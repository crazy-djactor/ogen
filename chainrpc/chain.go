package chainrpc

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"reflect"

	"github.com/olympus-protocol/ogen/chain"
	"github.com/olympus-protocol/ogen/chain/index"
	"github.com/olympus-protocol/ogen/primitives"
	"github.com/olympus-protocol/ogen/proto"
	"github.com/olympus-protocol/ogen/utils/chainhash"
)

type chainServer struct {
	chain *chain.Blockchain
	proto.UnimplementedChainServer
}

func (s *chainServer) GetChainInfo(ctx context.Context, _ *proto.Empty) (*proto.ChainInfo, error) {
	state := s.chain.State()
	validators := state.TipState().GetValidators()
	return &proto.ChainInfo{
		BlockHash:   state.Tip().Hash.String(),
		BlockHeight: state.Height(),
		Validators: &proto.ValidatorsInfo{
			Active:      validators.Active,
			PendingExit: validators.PendingExit,
			PenaltyExit: validators.PenaltyExit,
			Exited:      validators.Exited,
			Starting:    validators.Starting,
		},
	}, nil
}

func (s *chainServer) GetRawBlock(ctx context.Context, in *proto.Hash) (*proto.Block, error) {
	hash, err := chainhash.NewHashFromStr(in.Hash)
	if err != nil {
		return nil, err
	}
	block, err := s.chain.GetRawBlock(*hash)
	if err != nil {
		return nil, err
	}
	return &proto.Block{RawBlock: hex.EncodeToString(block)}, nil
}

func (s *chainServer) GetBlock(ctx context.Context, in *proto.Hash) (*proto.Block, error) {
	hash, err := chainhash.NewHashFromStr(in.Hash)
	if err != nil {
		return nil, err
	}
	block, err := s.chain.GetBlock(*hash)
	if err != nil {
		return nil, err
	}
	blockParse := &proto.Block{
		Hash: block.Hash().String(),
		Header: &proto.BlockHeader{
			Version:                    block.Header.Version,
			Nonce:                      block.Header.Nonce,
			TxMerkleRoot:               block.Header.TxMerkleRoot.String(),
			VoteMerkleRoot:             block.Header.VoteMerkleRoot.String(),
			DepositMerkleRoot:          block.Header.DepositMerkleRoot.String(),
			ExitMerkleRoot:             block.Header.ExitMerkleRoot.String(),
			VoteSlashingMerkleRoot:     block.Header.VoteSlashingMerkleRoot.String(),
			RandaoSlashingMerkleRoot:   block.Header.RANDAOSlashingMerkleRoot.String(),
			ProposerSlashingMerkleRoot: block.Header.ProposerSlashingMerkleRoot.String(),
			PrevBlockHash:              block.Header.PrevBlockHash.String(),
			Timestamp:                  block.Header.Timestamp,
			Slot:                       block.Header.Slot,
			StateRoot:                  block.Header.StateRoot.String(),
			FeeAddress:                 hex.EncodeToString(block.Header.FeeAddress[:]),
		},
		Txs:             block.GetTxs(),
		Signature:       hex.EncodeToString(block.Signature),
		RandaoSignature: hex.EncodeToString(block.RandaoSignature),
	}
	return blockParse, nil
}

func (s *chainServer) GetBlockHash(ctx context.Context, in *proto.Number) (*proto.Hash, error) {
	blockRow, exists := s.chain.State().Chain().GetNodeByHeight(in.Number)
	if !exists {
		return nil, errors.New("block not found")
	}
	return &proto.Hash{
		Hash: blockRow.Hash.String(),
	}, nil
}

func (s *chainServer) Sync(in *proto.Hash, stream proto.Chain_SyncServer) error {
	_, cancel := context.WithCancel(stream.Context())
	// Define starting point
	blockRow := new(index.BlockRow)
	defer cancel()
	// If user is on tip, silently close the channel
	if reflect.DeepEqual(in.Hash, s.chain.State().Tip().Hash.String()) {
		return nil
	}

	ok := true
	hash, err := chainhash.NewHashFromStr(in.Hash)
	if err != nil {
		return errors.New("unable to decode hash from string")
	}
	currBlockRow, ok := s.chain.State().GetRowByHash(*hash)
	if !ok {
		return errors.New("block starting point doesnt exist")
	}
	blockRow, ok = s.chain.State().Chain().Next(currBlockRow)
	if !ok {
		return errors.New("there is no next blockrow")
	}
	for {
		ok := true
		rawBlock, err := s.chain.GetRawBlock(blockRow.Hash)
		if err != nil {
			return errors.New("unable get raw block")
		}
		response := &proto.RawData{
			Data: hex.EncodeToString(rawBlock),
		}
		stream.Send(response)
		blockRow, ok = s.chain.State().Chain().Next(blockRow)
		if blockRow == nil || !ok {
			break
		}
	}
	return nil
}

type blockNotifee struct {
	blocks chan blockAndReceipts
}

type blockAndReceipts struct {
	block    *primitives.Block
	receipts []*primitives.EpochReceipt
	state    *primitives.State
}

func newBlockNotifee(ctx context.Context, chain *chain.Blockchain) blockNotifee {
	bn := blockNotifee{
		blocks: make(chan blockAndReceipts),
	}

	go func() {
		chain.Notify(&bn)

		<-ctx.Done()

		chain.Unnotify(&bn)
	}()

	return bn
}

func (bn *blockNotifee) NewTip(row *index.BlockRow, block *primitives.Block, newState *primitives.State, receipts []*primitives.EpochReceipt) {
	toSend := blockAndReceipts{block: block, receipts: receipts, state: newState}
	select {
	case bn.blocks <- toSend:
	default:
	}
}

func (bn *blockNotifee) ProposerSlashingConditionViolated(slashing primitives.ProposerSlashing) {}

func (s *chainServer) SubscribeBlocks(_ *proto.Empty, stream proto.Chain_SubscribeBlocksServer) error {
	bn := newBlockNotifee(stream.Context(), s.chain)

	for {
		select {
		case bl := <-bn.blocks:
			buf, err := bl.block.Marshal()
			if err != nil {
				return err
			}
			err = stream.Send(&proto.RawData{
				Data: hex.EncodeToString(buf),
			})
			if err != nil {
				return err
			}
		case <-stream.Context().Done():
			return nil
		}
	}
}

func (s *chainServer) SubscribeTransactions(in *proto.KeyPairs, stream proto.Chain_SubscribeTransactionsServer) error {
	bn := newBlockNotifee(stream.Context(), s.chain)
	accounts := make(map[[20]byte]struct{})
	for _, a := range in.Keys {
		account, err := hex.DecodeString(a)
		if err != nil {
			return err
		}
		if len(account) != 20 {
			return fmt.Errorf("expected public key hashes to be 20 bytes but got %d", len(a))
		}

		var acc [20]byte
		copy(acc[:], a)

		accounts[acc] = struct{}{}
	}

	for {
		select {
		case bl := <-bn.blocks:
			err := stream.Send(&proto.RawData{
				Data: hex.EncodeToString(bl.block.SerializedTx(accounts)),
			})
			if err != nil {
				return err
			}
		case <-stream.Context().Done():
			return nil
		}
	}
}

func (s *chainServer) SubscribeValidatorTransaction(in *proto.KeyPairs, stream proto.Chain_SubscribeValidatorTransactionsServer) error {
	bn := newBlockNotifee(stream.Context(), s.chain)
	accounts := make(map[[48]byte]struct{})
	for _, a := range in.Keys {
		pubkey, err := hex.DecodeString(a)
		if err != nil {
			return err
		}
		if len(pubkey) != 48 {
			return fmt.Errorf("expected public key to be 48 bytes but got %d", len(a))
		}

		var acc [48]byte
		copy(acc[:], a)

		accounts[acc] = struct{}{}
	}

	for {
		select {
		case bl := <-bn.blocks:
			err := stream.Send(&proto.RawData{
				Data: hex.EncodeToString(bl.block.SerializedEpochs(accounts)),
			})
			if err != nil {
				return err
			}
		case <-stream.Context().Done():
			return nil
		}
	}
}

func (s *chainServer) GetAccountInfo(ctx context.Context, data *proto.Account) (*proto.AccountInfo, error) {
	var account [20]byte
	accBytes, err := hex.DecodeString(data.Account)
	if err != nil {
		return nil, err
	}
	copy(account[:], accBytes)

	balance := s.chain.State().TipState().CoinsState.Balances[account]
	nonce := s.chain.State().TipState().CoinsState.Nonces[account]
	txs, err := s.chain.GetAccountTxs(account)
	if err != nil {
		return nil, err
	}

	accInfo := &proto.AccountInfo{
		Account: data.Account,
		Txs:     txs.Strings(),
		Balance: balance,
		Nonce:   nonce,
	}
	return accInfo, nil
}

func (s *chainServer) GetTransaction(ctx context.Context, h *proto.Hash) (*proto.Tx, error) {
	txid, err := chainhash.NewHashFromStr(h.Hash)
	if err != nil {
		return nil, err
	}
	tx, err := s.chain.GetTx(*txid)
	if err != nil {
		return nil, err
	}
	payload, err := tx.GetPayload()
	if err != nil {
		return nil, err
	}
	txParse := &proto.Tx{
		Hash:    tx.Hash().String(),
		Version: tx.Version,
		Type:    tx.Type,
	}
	switch p := payload.(type) {
	case *primitives.TransferSinglePayload:
		txParse.TransferSinglePayload = &proto.TransferSingle{
			To:            hex.EncodeToString(p.To[:]),
			FromPublicKey: hex.EncodeToString(p.FromPublicKey[:]),
			Amount:        p.Amount,
			Nonce:         p.Nonce,
			Fee:           p.Fee,
			Signature:     hex.EncodeToString(p.Signature),
		}
	case *primitives.TransferMultiPayload:
		txParse.TransferMultiPayload = &proto.TransferMulti{
			To:        hex.EncodeToString(p.To[:]),
			Amount:    p.Amount,
			Nonce:     p.Nonce,
			Fee:       p.Fee,
			Signature: hex.EncodeToString(p.MultiSig),
		}
	}

	return txParse, nil
}

var _ proto.ChainServer = &chainServer{}

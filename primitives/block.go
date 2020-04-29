package primitives

import (
	"io"
	"time"

	"github.com/olympus-protocol/ogen/bls"
	"github.com/olympus-protocol/ogen/utils/chainhash"
	"github.com/olympus-protocol/ogen/utils/serializer"
)

const (
	maxBlockSize = 1024 * 512 // 512 kilobytes
)

// Block is a block in the blockchain.
type Block struct {
	Header          BlockHeader
	Votes           []MultiValidatorVote
	Txs             []Tx
	Signature       [96]byte
	RandaoSignature [96]byte
}

func (b *Block) MinerSig() (*bls.Signature, error) {
	return bls.DeserializeSignature(b.Signature)
}

func (b *Block) GetTime() time.Time {
	return b.Header.Timestamp
}

func (b *Block) GetTx(index int32) *Tx {
	return &b.Txs[index]
}

func (b *Block) Hash() chainhash.Hash {
	return b.Header.Hash()
}

func (m *Block) Encode(w io.Writer) error {
	err := m.Header.Serialize(w)
	if err != nil {
		return err
	}
	err = serializer.WriteVarInt(w, uint64(len(m.Txs)))
	if err != nil {
		return err
	}
	for _, tx := range m.Txs {
		err := tx.Encode(w)
		if err != nil {
			return err
		}
	}
	err = serializer.WriteVarInt(w, uint64(len(m.Votes)))
	if err != nil {
		return err
	}
	for _, vote := range m.Votes {
		err := vote.Serialize(w)
		if err != nil {
			return err
		}
	}
	err = serializer.WriteElements(w, m.Signature, m.RandaoSignature)
	if err != nil {
		return err
	}
	return nil
}

func (m *Block) Decode(r io.Reader) error {
	err := m.Header.Deserialize(r)
	if err != nil {
		return err
	}
	txCount, err := serializer.ReadVarInt(r)
	if err != nil {
		return err
	}
	m.Txs = make([]Tx, txCount)
	for i := uint64(0); i < txCount; i++ {
		err := m.Txs[i].Decode(r)
		if err != nil {
			return err
		}
	}
	voteCount, err := serializer.ReadVarInt(r)
	if err != nil {
		return err
	}
	m.Votes = make([]MultiValidatorVote, voteCount)
	for i := range m.Votes {
		err := m.Votes[i].Deserialize(r)
		if err != nil {
			return err
		}
	}
	err = serializer.ReadElements(r, &m.Signature, &m.RandaoSignature)
	if err != nil {
		return err
	}
	return nil
}

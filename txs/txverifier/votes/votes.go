package votes_txverifier

import (
	"bytes"
	"errors"
	"sync"

	"github.com/olympus-protocol/ogen/bls"
	"github.com/olympus-protocol/ogen/params"
	"github.com/olympus-protocol/ogen/primitives"
	"github.com/olympus-protocol/ogen/txs/txpayloads"
	votes_txpayload "github.com/olympus-protocol/ogen/txs/txpayloads/votes"
)

var (
	ErrorNoTypeSpecified  = errors.New("voteTx-no-type-rules: the selected action is not specified on the tx verifier scheme")
	ErrorInvalidSignature = errors.New("voteTx-invalid-signature: the signature verification is invalid")
	ErrorMatchDataNoExist = errors.New("voteTx-not-found-match-data: the match data searched doesn't exists")
	ErrorDataNoMatch      = errors.New("voteTx-invalid-match-data: the data used to sign the transaction doesn't match")
)

type VotesTxVerifier struct {
	state  *primitives.State
	params *params.ChainParams
}

func (v VotesTxVerifier) DeserializePayload(payload []byte, Action primitives.TxAction) (txpayloads.Payload, error) {
	var Payload txpayloads.Payload
	switch Action {
	case primitives.Upload:
		Payload = new(votes_txpayload.PayloadUploadAndUpdate)
	case primitives.Update:
		Payload = new(votes_txpayload.PayloadUploadAndUpdate)
	case primitives.Revoke:
		Payload = new(votes_txpayload.PayloadRevoke)
	default:
		return nil, ErrorNoTypeSpecified
	}
	buf := bytes.NewBuffer(payload)
	err := Payload.Deserialize(buf)
	if err != nil {
		return nil, err
	}
	return Payload, nil
}

func (v VotesTxVerifier) SigVerify(payload []byte, Action primitives.TxAction) error {
	VerPayload, err := v.DeserializePayload(payload, Action)
	if err != nil {
		return err
	}
	pubKey, err := VerPayload.GetPublicKey()
	if err != nil {
		return err
	}
	msg, err := VerPayload.GetMessage()
	if err != nil {
		return err
	}
	sig, err := VerPayload.GetSignature()
	if err != nil {
		return err
	}
	valid, err := bls.VerifySig(pubKey, msg, sig)
	if err != nil {
		return err
	}
	if !valid {
		return ErrorInvalidSignature
	}
	return nil
}

type routineRes struct {
	Err error
}

func (v VotesTxVerifier) SigVerifyBatch(payload [][]byte, Action primitives.TxAction) error {
	var wg sync.WaitGroup
	doneChan := make(chan routineRes, len(payload))
	for _, singlePayload := range payload {
		wg.Add(1)
		go func(wg *sync.WaitGroup, payload []byte) {
			defer wg.Done()
			var resp routineRes
			err := v.SigVerify(payload, Action)
			if err != nil {
				resp.Err = err
			}
			doneChan <- resp
			return
		}(&wg, singlePayload)
	}
	wg.Wait()
	doneRes := <-doneChan
	if doneRes.Err != nil {
		return doneRes.Err
	}
	return nil
}

func (v VotesTxVerifier) MatchVerify(payload []byte, Action primitives.TxAction) error {
	// VerPayload, err := v.DeserializePayload(payload, Action)
	// if err != nil {
	// 	return err
	// }
	// searchHash, err := VerPayload.GetHashForDataMatch()
	// if err != nil {
	// 	return err
	// }
	// ok := v.state.WorkerState.Have(searchHash)
	// if !ok {
	// 	return ErrorMatchDataNoExist
	// }
	// data := v.state.WorkerState.Get(searchHash)
	// pubKey, err := bls.DeserializePublicKey(data.PubKey)
	// if err != nil {
	// 	return err
	// }
	// matchPubKey, err := VerPayload.GetPublicKey()
	// if err != nil {
	// 	return err
	// }
	// equal := reflect.DeepEqual(pubKey, matchPubKey)
	// if !equal {
	// 	return ErrorDataNoMatch
	// }
	return nil
}

func (v VotesTxVerifier) MatchVerifyBatch(payload [][]byte, Action primitives.TxAction) error {
	var wg sync.WaitGroup
	doneChan := make(chan routineRes, len(payload))
	for _, singlePayload := range payload {
		wg.Add(1)
		go func(wg *sync.WaitGroup, payload []byte) {
			defer wg.Done()
			var resp routineRes
			err := v.MatchVerify(payload, Action)
			if err != nil {
				resp.Err = err
			}
			doneChan <- resp
			return
		}(&wg, singlePayload)
	}
	wg.Wait()
	doneRes := <-doneChan
	if doneRes.Err != nil {
		return doneRes.Err
	}
	return nil
}

func NewVotesTxVerifier(state *primitives.State, params *params.ChainParams) VotesTxVerifier {
	v := VotesTxVerifier{
		state:  state,
		params: params,
	}
	return v
}

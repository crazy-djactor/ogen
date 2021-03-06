// Code generated by MockGen. DO NOT EDIT.
// Source: internal/chain/state.go

// Package chain is a generated GoMock package.
package chain

import (
	gomock "github.com/golang/mock/gomock"
	chainindex "github.com/olympus-protocol/ogen/internal/chainindex"
	state "github.com/olympus-protocol/ogen/internal/state"
	chainhash "github.com/olympus-protocol/ogen/pkg/chainhash"
	params "github.com/olympus-protocol/ogen/pkg/params"
	primitives "github.com/olympus-protocol/ogen/pkg/primitives"
	reflect "reflect"
)

// MockStateService is a mock of StateService interface
type MockStateService struct {
	ctrl     *gomock.Controller
	recorder *MockStateServiceMockRecorder
}

// MockStateServiceMockRecorder is the mock recorder for MockStateService
type MockStateServiceMockRecorder struct {
	mock *MockStateService
}

// NewMockStateService creates a new mock instance
func NewMockStateService(ctrl *gomock.Controller) *MockStateService {
	mock := &MockStateService{ctrl: ctrl}
	mock.recorder = &MockStateServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStateService) EXPECT() *MockStateServiceMockRecorder {
	return m.recorder
}

// Blockchain mocks base method
func (m *MockStateService) Blockchain() *Chain {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Blockchain")
	ret0, _ := ret[0].(*Chain)
	return ret0
}

// Blockchain indicates an expected call of Blockchain
func (mr *MockStateServiceMockRecorder) Blockchain() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Blockchain", reflect.TypeOf((*MockStateService)(nil).Blockchain))
}

// GetLatestVote mocks base method
func (m *MockStateService) GetLatestVote(val uint64) (*primitives.MultiValidatorVote, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestVote", val)
	ret0, _ := ret[0].(*primitives.MultiValidatorVote)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetLatestVote indicates an expected call of GetLatestVote
func (mr *MockStateServiceMockRecorder) GetLatestVote(val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestVote", reflect.TypeOf((*MockStateService)(nil).GetLatestVote), val)
}

// SetLatestVotesIfNeeded mocks base method
func (m *MockStateService) SetLatestVotesIfNeeded(vals []uint64, vote *primitives.MultiValidatorVote) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLatestVotesIfNeeded", vals, vote)
}

// SetLatestVotesIfNeeded indicates an expected call of SetLatestVotesIfNeeded
func (mr *MockStateServiceMockRecorder) SetLatestVotesIfNeeded(vals, vote interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLatestVotesIfNeeded", reflect.TypeOf((*MockStateService)(nil).SetLatestVotesIfNeeded), vals, vote)
}

// Chain mocks base method
func (m *MockStateService) Chain() *Chain {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Chain")
	ret0, _ := ret[0].(*Chain)
	return ret0
}

// Chain indicates an expected call of Chain
func (mr *MockStateServiceMockRecorder) Chain() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Chain", reflect.TypeOf((*MockStateService)(nil).Chain))
}

// Index mocks base method
func (m *MockStateService) Index() *chainindex.BlockIndex {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Index")
	ret0, _ := ret[0].(*chainindex.BlockIndex)
	return ret0
}

// Index indicates an expected call of Index
func (mr *MockStateServiceMockRecorder) Index() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Index", reflect.TypeOf((*MockStateService)(nil).Index))
}

// SetFinalizedHead mocks base method
func (m *MockStateService) SetFinalizedHead(finalizedHash chainhash.Hash, finalizedState state.State) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetFinalizedHead", finalizedHash, finalizedState)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetFinalizedHead indicates an expected call of SetFinalizedHead
func (mr *MockStateServiceMockRecorder) SetFinalizedHead(finalizedHash, finalizedState interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFinalizedHead", reflect.TypeOf((*MockStateService)(nil).SetFinalizedHead), finalizedHash, finalizedState)
}

// GetFinalizedHead mocks base method
func (m *MockStateService) GetFinalizedHead() (*chainindex.BlockRow, state.State) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFinalizedHead")
	ret0, _ := ret[0].(*chainindex.BlockRow)
	ret1, _ := ret[1].(state.State)
	return ret0, ret1
}

// GetFinalizedHead indicates an expected call of GetFinalizedHead
func (mr *MockStateServiceMockRecorder) GetFinalizedHead() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFinalizedHead", reflect.TypeOf((*MockStateService)(nil).GetFinalizedHead))
}

// SetJustifiedHead mocks base method
func (m *MockStateService) SetJustifiedHead(justifiedHash chainhash.Hash, justifiedState state.State) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetJustifiedHead", justifiedHash, justifiedState)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetJustifiedHead indicates an expected call of SetJustifiedHead
func (mr *MockStateServiceMockRecorder) SetJustifiedHead(justifiedHash, justifiedState interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetJustifiedHead", reflect.TypeOf((*MockStateService)(nil).SetJustifiedHead), justifiedHash, justifiedState)
}

// GetJustifiedHead mocks base method
func (m *MockStateService) GetJustifiedHead() (*chainindex.BlockRow, state.State) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJustifiedHead")
	ret0, _ := ret[0].(*chainindex.BlockRow)
	ret1, _ := ret[1].(state.State)
	return ret0, ret1
}

// GetJustifiedHead indicates an expected call of GetJustifiedHead
func (mr *MockStateServiceMockRecorder) GetJustifiedHead() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJustifiedHead", reflect.TypeOf((*MockStateService)(nil).GetJustifiedHead))
}

// GetStateForHash mocks base method
func (m *MockStateService) GetStateForHash(hash chainhash.Hash) (state.State, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStateForHash", hash)
	ret0, _ := ret[0].(state.State)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetStateForHash indicates an expected call of GetStateForHash
func (mr *MockStateServiceMockRecorder) GetStateForHash(hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStateForHash", reflect.TypeOf((*MockStateService)(nil).GetStateForHash), hash)
}

// GetStateForHashAtSlot mocks base method
func (m *MockStateService) GetStateForHashAtSlot(hash chainhash.Hash, slot uint64, view state.BlockView, p *params.ChainParams) (state.State, []*primitives.EpochReceipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStateForHashAtSlot", hash, slot, view, p)
	ret0, _ := ret[0].(state.State)
	ret1, _ := ret[1].([]*primitives.EpochReceipt)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetStateForHashAtSlot indicates an expected call of GetStateForHashAtSlot
func (mr *MockStateServiceMockRecorder) GetStateForHashAtSlot(hash, slot, view, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStateForHashAtSlot", reflect.TypeOf((*MockStateService)(nil).GetStateForHashAtSlot), hash, slot, view, p)
}

// Add mocks base method
func (m *MockStateService) Add(block *primitives.Block) (state.State, []*primitives.EpochReceipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", block)
	ret0, _ := ret[0].(state.State)
	ret1, _ := ret[1].([]*primitives.EpochReceipt)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Add indicates an expected call of Add
func (mr *MockStateServiceMockRecorder) Add(block interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockStateService)(nil).Add), block)
}

// RemoveBeforeSlot mocks base method
func (m *MockStateService) RemoveBeforeSlot(slot uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveBeforeSlot", slot)
}

// RemoveBeforeSlot indicates an expected call of RemoveBeforeSlot
func (mr *MockStateServiceMockRecorder) RemoveBeforeSlot(slot interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveBeforeSlot", reflect.TypeOf((*MockStateService)(nil).RemoveBeforeSlot), slot)
}

// GetRowByHash mocks base method
func (m *MockStateService) GetRowByHash(h chainhash.Hash) (*chainindex.BlockRow, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRowByHash", h)
	ret0, _ := ret[0].(*chainindex.BlockRow)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetRowByHash indicates an expected call of GetRowByHash
func (mr *MockStateServiceMockRecorder) GetRowByHash(h interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRowByHash", reflect.TypeOf((*MockStateService)(nil).GetRowByHash), h)
}

// Height mocks base method
func (m *MockStateService) Height() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Height")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Height indicates an expected call of Height
func (mr *MockStateServiceMockRecorder) Height() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Height", reflect.TypeOf((*MockStateService)(nil).Height))
}

// TipState mocks base method
func (m *MockStateService) TipState() state.State {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TipState")
	ret0, _ := ret[0].(state.State)
	return ret0
}

// TipState indicates an expected call of TipState
func (mr *MockStateServiceMockRecorder) TipState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TipState", reflect.TypeOf((*MockStateService)(nil).TipState))
}

// TipStateAtSlot mocks base method
func (m *MockStateService) TipStateAtSlot(slot uint64) (state.State, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TipStateAtSlot", slot)
	ret0, _ := ret[0].(state.State)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TipStateAtSlot indicates an expected call of TipStateAtSlot
func (mr *MockStateServiceMockRecorder) TipStateAtSlot(slot interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TipStateAtSlot", reflect.TypeOf((*MockStateService)(nil).TipStateAtSlot), slot)
}

// GetSubView mocks base method
func (m *MockStateService) GetSubView(tip chainhash.Hash) (View, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubView", tip)
	ret0, _ := ret[0].(View)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubView indicates an expected call of GetSubView
func (mr *MockStateServiceMockRecorder) GetSubView(tip interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubView", reflect.TypeOf((*MockStateService)(nil).GetSubView), tip)
}

// Tip mocks base method
func (m *MockStateService) Tip() *chainindex.BlockRow {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tip")
	ret0, _ := ret[0].(*chainindex.BlockRow)
	return ret0
}

// Tip indicates an expected call of Tip
func (mr *MockStateServiceMockRecorder) Tip() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tip", reflect.TypeOf((*MockStateService)(nil).Tip))
}

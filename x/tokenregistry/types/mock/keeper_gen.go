// Code generated by MockGen. DO NOT EDIT.
// Source: ../expected_keepers.go

// Package tokenregistrymocks is a generated GoMock package.
package tokenregistrymocks

import (
	reflect "reflect"

	types "github.com/Sifchain/sifnode/x/tokenregistry/types"
	types0 "github.com/cosmos/cosmos-sdk/types"
	gomock "github.com/golang/mock/gomock"
	types1 "github.com/tendermint/tendermint/abci/types"
)

// MockKeeper is a mock of Keeper interface.
type MockKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockKeeperMockRecorder
}

// MockKeeperMockRecorder is the mock recorder for MockKeeper.
type MockKeeperMockRecorder struct {
	mock *MockKeeper
}

// NewMockKeeper creates a new mock instance.
func NewMockKeeper(ctrl *gomock.Controller) *MockKeeper {
	mock := &MockKeeper{ctrl: ctrl}
	mock.recorder = &MockKeeperMockRecorder{mock}
	return mock
}

func (m *MockKeeper) CheckDenomPermissions(ctx types0.Context, denom string, permissions []types.Permission) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckDenomPermissions", ctx, denom)
	ret0, _ := ret[0].(bool)
	return ret0
}
func (mr *MockKeeperMockRecorder) CheckDenomPermissions(ctx, denom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckDenomPermissions", reflect.TypeOf((*MockKeeper)(nil).CheckDenomPermissions), ctx, denom)
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKeeper) EXPECT() *MockKeeperMockRecorder {
	return m.recorder
}

// ExportGenesis mocks base method.
func (m *MockKeeper) ExportGenesis(ctx types0.Context) *types.GenesisState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportGenesis", ctx)
	ret0, _ := ret[0].(*types.GenesisState)
	return ret0
}

// ExportGenesis indicates an expected call of ExportGenesis.
func (mr *MockKeeperMockRecorder) ExportGenesis(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportGenesis", reflect.TypeOf((*MockKeeper)(nil).ExportGenesis), ctx)
}

// GetDenom mocks base method.
func (m *MockKeeper) GetDenom(ctx types0.Context, denom string) types.RegistryEntry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDenom", ctx, denom)
	ret0, _ := ret[0].(types.RegistryEntry)
	return ret0
}

// GetDenom indicates an expected call of GetDenom.
func (mr *MockKeeperMockRecorder) GetDenom(ctx, denom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDenom", reflect.TypeOf((*MockKeeper)(nil).GetDenom), ctx, denom)
}

// GetDenomWhitelist mocks base method.
func (m *MockKeeper) GetDenomWhitelist(ctx types0.Context) types.Registry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDenomWhitelist", ctx)
	ret0, _ := ret[0].(types.Registry)
	return ret0
}

// GetDenomWhitelist indicates an expected call of GetDenomWhitelist.
func (mr *MockKeeperMockRecorder) GetDenomWhitelist(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDenomWhitelist", reflect.TypeOf((*MockKeeper)(nil).GetDenomWhitelist), ctx)
}

// InitGenesis mocks base method.
func (m *MockKeeper) InitGenesis(ctx types0.Context, state types.GenesisState) []types1.ValidatorUpdate {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitGenesis", ctx, state)
	ret0, _ := ret[0].([]types1.ValidatorUpdate)
	return ret0
}

// InitGenesis indicates an expected call of InitGenesis.
func (mr *MockKeeperMockRecorder) InitGenesis(ctx, state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitGenesis", reflect.TypeOf((*MockKeeper)(nil).InitGenesis), ctx, state)
}

// IsAdminAccount mocks base method.
func (m *MockKeeper) IsAdminAccount(ctx types0.Context, adminAccount types0.AccAddress) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAdminAccount", ctx, adminAccount)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsAdminAccount indicates an expected call of IsAdminAccount.
func (mr *MockKeeperMockRecorder) IsAdminAccount(ctx, adminAccount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAdminAccount", reflect.TypeOf((*MockKeeper)(nil).IsAdminAccount), ctx, adminAccount)
}

// IsDenomWhitelisted mocks base method.
func (m *MockKeeper) IsDenomWhitelisted(ctx types0.Context, denom string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDenomWhitelisted", ctx, denom)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsDenomWhitelisted indicates an expected call of IsDenomWhitelisted.
func (mr *MockKeeperMockRecorder) IsDenomWhitelisted(ctx, denom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDenomWhitelisted", reflect.TypeOf((*MockKeeper)(nil).IsDenomWhitelisted), ctx, denom)
}

// RemoveToken mocks base method.
func (m *MockKeeper) RemoveToken(ctx types0.Context, denom string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveToken", ctx, denom)
}

// RemoveToken indicates an expected call of RemoveToken.
func (mr *MockKeeperMockRecorder) RemoveToken(ctx, denom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveToken", reflect.TypeOf((*MockKeeper)(nil).RemoveToken), ctx, denom)
}

// SetAdminAccount mocks base method.
func (m *MockKeeper) SetAdminAccount(ctx types0.Context, adminAccount types0.AccAddress) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAdminAccount", ctx, adminAccount)
}

// SetAdminAccount indicates an expected call of SetAdminAccount.
func (mr *MockKeeperMockRecorder) SetAdminAccount(ctx, adminAccount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAdminAccount", reflect.TypeOf((*MockKeeper)(nil).SetAdminAccount), ctx, adminAccount)
}

// SetToken mocks base method.
func (m *MockKeeper) SetToken(ctx types0.Context, entry *types.RegistryEntry) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetToken", ctx, entry)
}

// SetToken indicates an expected call of SetToken.
func (mr *MockKeeperMockRecorder) SetToken(ctx, entry interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetToken", reflect.TypeOf((*MockKeeper)(nil).SetToken), ctx, entry)
}

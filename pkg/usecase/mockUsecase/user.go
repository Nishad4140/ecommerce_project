// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/usecase/interface/user.go

// Package mockUsecase is a generated GoMock package.
package mockUsecase

import (
	reflect "reflect"

	helper "github.com/Nishad4140/ecommerce_project/pkg/common/helperStruct"
	response "github.com/Nishad4140/ecommerce_project/pkg/common/response"
	gomock "github.com/golang/mock/gomock"
)

// MockUserUseCase is a mock of UserUseCase interface.
type MockUserUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUserUseCaseMockRecorder
}

// MockUserUseCaseMockRecorder is the mock recorder for MockUserUseCase.
type MockUserUseCaseMockRecorder struct {
	mock *MockUserUseCase
}

// NewMockUserUseCase creates a new mock instance.
func NewMockUserUseCase(ctrl *gomock.Controller) *MockUserUseCase {
	mock := &MockUserUseCase{ctrl: ctrl}
	mock.recorder = &MockUserUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUseCase) EXPECT() *MockUserUseCaseMockRecorder {
	return m.recorder
}

// AddAddress mocks base method.
func (m *MockUserUseCase) AddAddress(id int, address helper.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAddress", id, address)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAddress indicates an expected call of AddAddress.
func (mr *MockUserUseCaseMockRecorder) AddAddress(id, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAddress", reflect.TypeOf((*MockUserUseCase)(nil).AddAddress), id, address)
}

// CreateWallet mocks base method.
func (m *MockUserUseCase) CreateWallet(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWallet", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateWallet indicates an expected call of CreateWallet.
func (mr *MockUserUseCaseMockRecorder) CreateWallet(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWallet", reflect.TypeOf((*MockUserUseCase)(nil).CreateWallet), id)
}

// EditProfile mocks base method.
func (m *MockUserUseCase) EditProfile(userID int, updatingDetails helper.UpdateProfile) (response.Userprofile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditProfile", userID, updatingDetails)
	ret0, _ := ret[0].(response.Userprofile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditProfile indicates an expected call of EditProfile.
func (mr *MockUserUseCaseMockRecorder) EditProfile(userID, updatingDetails interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditProfile", reflect.TypeOf((*MockUserUseCase)(nil).EditProfile), userID, updatingDetails)
}

// ForgotPassword mocks base method.
func (m *MockUserUseCase) ForgotPassword(forgotPass helper.ForgotPassword) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForgotPassword", forgotPass)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForgotPassword indicates an expected call of ForgotPassword.
func (mr *MockUserUseCaseMockRecorder) ForgotPassword(forgotPass interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForgotPassword", reflect.TypeOf((*MockUserUseCase)(nil).ForgotPassword), forgotPass)
}

// UpdateAddress mocks base method.
func (m *MockUserUseCase) UpdateAddress(id, addressId int, address helper.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAddress", id, addressId, address)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAddress indicates an expected call of UpdateAddress.
func (mr *MockUserUseCaseMockRecorder) UpdateAddress(id, addressId, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAddress", reflect.TypeOf((*MockUserUseCase)(nil).UpdateAddress), id, addressId, address)
}

// UpdatePassword mocks base method.
func (m *MockUserUseCase) UpdatePassword(userID int, Passwords helper.UpdatePassword) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", userID, Passwords)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePassword indicates an expected call of UpdatePassword.
func (mr *MockUserUseCaseMockRecorder) UpdatePassword(userID, Passwords interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockUserUseCase)(nil).UpdatePassword), userID, Passwords)
}

// UserLogin mocks base method.
func (m *MockUserUseCase) UserLogin(user helper.LoginReq) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserLogin", user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserLogin indicates an expected call of UserLogin.
func (mr *MockUserUseCaseMockRecorder) UserLogin(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserLogin", reflect.TypeOf((*MockUserUseCase)(nil).UserLogin), user)
}

// UserSignUp mocks base method.
func (m *MockUserUseCase) UserSignUp(user helper.UserReq) (response.UserData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserSignUp", user)
	ret0, _ := ret[0].(response.UserData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserSignUp indicates an expected call of UserSignUp.
func (mr *MockUserUseCaseMockRecorder) UserSignUp(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserSignUp", reflect.TypeOf((*MockUserUseCase)(nil).UserSignUp), user)
}

// VerifyWallet mocks base method.
func (m *MockUserUseCase) VerifyWallet(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyWallet", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyWallet indicates an expected call of VerifyWallet.
func (mr *MockUserUseCaseMockRecorder) VerifyWallet(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyWallet", reflect.TypeOf((*MockUserUseCase)(nil).VerifyWallet), id)
}

// ViewProfile mocks base method.
func (m *MockUserUseCase) ViewProfile(userID int) (response.Userprofile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewProfile", userID)
	ret0, _ := ret[0].(response.Userprofile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ViewProfile indicates an expected call of ViewProfile.
func (mr *MockUserUseCaseMockRecorder) ViewProfile(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewProfile", reflect.TypeOf((*MockUserUseCase)(nil).ViewProfile), userID)
}

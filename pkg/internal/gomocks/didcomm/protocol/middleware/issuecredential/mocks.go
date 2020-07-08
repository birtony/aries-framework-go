// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hyperledger/aries-framework-go/pkg/didcomm/protocol/middleware/issuecredential (interfaces: Provider,Metadata)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	service "github.com/hyperledger/aries-framework-go/pkg/didcomm/common/service"
	issuecredential "github.com/hyperledger/aries-framework-go/pkg/didcomm/protocol/issuecredential"
	vdri "github.com/hyperledger/aries-framework-go/pkg/framework/aries/api/vdri"
	verifiable "github.com/hyperledger/aries-framework-go/pkg/store/verifiable"
	reflect "reflect"
)

// MockProvider is a mock of Provider interface
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// VDRIRegistry mocks base method
func (m *MockProvider) VDRIRegistry() vdri.Registry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VDRIRegistry")
	ret0, _ := ret[0].(vdri.Registry)
	return ret0
}

// VDRIRegistry indicates an expected call of VDRIRegistry
func (mr *MockProviderMockRecorder) VDRIRegistry() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VDRIRegistry", reflect.TypeOf((*MockProvider)(nil).VDRIRegistry))
}

// VerifiableStore mocks base method
func (m *MockProvider) VerifiableStore() verifiable.Store {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifiableStore")
	ret0, _ := ret[0].(verifiable.Store)
	return ret0
}

// VerifiableStore indicates an expected call of VerifiableStore
func (mr *MockProviderMockRecorder) VerifiableStore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifiableStore", reflect.TypeOf((*MockProvider)(nil).VerifiableStore))
}

// MockMetadata is a mock of Metadata interface
type MockMetadata struct {
	ctrl     *gomock.Controller
	recorder *MockMetadataMockRecorder
}

// MockMetadataMockRecorder is the mock recorder for MockMetadata
type MockMetadataMockRecorder struct {
	mock *MockMetadata
}

// NewMockMetadata creates a new mock instance
func NewMockMetadata(ctrl *gomock.Controller) *MockMetadata {
	mock := &MockMetadata{ctrl: ctrl}
	mock.recorder = &MockMetadataMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMetadata) EXPECT() *MockMetadataMockRecorder {
	return m.recorder
}

// CredentialNames mocks base method
func (m *MockMetadata) CredentialNames() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CredentialNames")
	ret0, _ := ret[0].([]string)
	return ret0
}

// CredentialNames indicates an expected call of CredentialNames
func (mr *MockMetadataMockRecorder) CredentialNames() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CredentialNames", reflect.TypeOf((*MockMetadata)(nil).CredentialNames))
}

// IssueCredential mocks base method
func (m *MockMetadata) IssueCredential() *issuecredential.IssueCredential {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IssueCredential")
	ret0, _ := ret[0].(*issuecredential.IssueCredential)
	return ret0
}

// IssueCredential indicates an expected call of IssueCredential
func (mr *MockMetadataMockRecorder) IssueCredential() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IssueCredential", reflect.TypeOf((*MockMetadata)(nil).IssueCredential))
}

// Message mocks base method
func (m *MockMetadata) Message() service.DIDCommMsg {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Message")
	ret0, _ := ret[0].(service.DIDCommMsg)
	return ret0
}

// Message indicates an expected call of Message
func (mr *MockMetadataMockRecorder) Message() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Message", reflect.TypeOf((*MockMetadata)(nil).Message))
}

// OfferCredential mocks base method
func (m *MockMetadata) OfferCredential() *issuecredential.OfferCredential {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OfferCredential")
	ret0, _ := ret[0].(*issuecredential.OfferCredential)
	return ret0
}

// OfferCredential indicates an expected call of OfferCredential
func (mr *MockMetadataMockRecorder) OfferCredential() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OfferCredential", reflect.TypeOf((*MockMetadata)(nil).OfferCredential))
}

// Properties mocks base method
func (m *MockMetadata) Properties() map[string]interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Properties")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// Properties indicates an expected call of Properties
func (mr *MockMetadataMockRecorder) Properties() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Properties", reflect.TypeOf((*MockMetadata)(nil).Properties))
}

// ProposeCredential mocks base method
func (m *MockMetadata) ProposeCredential() *issuecredential.ProposeCredential {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProposeCredential")
	ret0, _ := ret[0].(*issuecredential.ProposeCredential)
	return ret0
}

// ProposeCredential indicates an expected call of ProposeCredential
func (mr *MockMetadataMockRecorder) ProposeCredential() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProposeCredential", reflect.TypeOf((*MockMetadata)(nil).ProposeCredential))
}

// StateName mocks base method
func (m *MockMetadata) StateName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateName")
	ret0, _ := ret[0].(string)
	return ret0
}

// StateName indicates an expected call of StateName
func (mr *MockMetadataMockRecorder) StateName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateName", reflect.TypeOf((*MockMetadata)(nil).StateName))
}

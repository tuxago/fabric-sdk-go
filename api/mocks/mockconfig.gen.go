// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hyperledger/fabric-sdk-go/api (interfaces: Config)

package mock_api

import (
	x509 "crypto/x509"
	gomock "github.com/golang/mock/gomock"
	api "github.com/tuxago/fabric-sdk-go/api"
	factory "github.com/hyperledger/fabric/bccsp/factory"
	viper "github.com/spf13/viper"
)

// MockConfig is a mock of Config interface
type MockConfig struct {
	ctrl     *gomock.Controller
	recorder *MockConfigMockRecorder
}

// MockConfigMockRecorder is the mock recorder for MockConfig
type MockConfigMockRecorder struct {
	mock *MockConfig
}

// NewMockConfig creates a new mock instance
func NewMockConfig(ctrl *gomock.Controller) *MockConfig {
	mock := &MockConfig{ctrl: ctrl}
	mock.recorder = &MockConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockConfig) EXPECT() *MockConfigMockRecorder {
	return _m.recorder
}

// GetCSPConfig mocks base method
func (_m *MockConfig) GetCSPConfig() *factory.FactoryOpts {
	ret := _m.ctrl.Call(_m, "GetCSPConfig")
	ret0, _ := ret[0].(*factory.FactoryOpts)
	return ret0
}

// GetCSPConfig indicates an expected call of GetCSPConfig
func (_mr *MockConfigMockRecorder) GetCSPConfig() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetCSPConfig")
}

// GetCryptoConfigPath mocks base method
func (_m *MockConfig) GetCryptoConfigPath() string {
	ret := _m.ctrl.Call(_m, "GetCryptoConfigPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetCryptoConfigPath indicates an expected call of GetCryptoConfigPath
func (_mr *MockConfigMockRecorder) GetCryptoConfigPath() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetCryptoConfigPath")
}

// GetFabricCAClientCertFile mocks base method
func (_m *MockConfig) GetFabricCAClientCertFile() string {
	ret := _m.ctrl.Call(_m, "GetFabricCAClientCertFile")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetFabricCAClientCertFile indicates an expected call of GetFabricCAClientCertFile
func (_mr *MockConfigMockRecorder) GetFabricCAClientCertFile() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetFabricCAClientCertFile")
}

// GetFabricCAClientKeyFile mocks base method
func (_m *MockConfig) GetFabricCAClientKeyFile() string {
	ret := _m.ctrl.Call(_m, "GetFabricCAClientKeyFile")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetFabricCAClientKeyFile indicates an expected call of GetFabricCAClientKeyFile
func (_mr *MockConfigMockRecorder) GetFabricCAClientKeyFile() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetFabricCAClientKeyFile")
}

// GetFabricCAHomeDir mocks base method
func (_m *MockConfig) GetFabricCAHomeDir() string {
	ret := _m.ctrl.Call(_m, "GetFabricCAHomeDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetFabricCAHomeDir indicates an expected call of GetFabricCAHomeDir
func (_mr *MockConfigMockRecorder) GetFabricCAHomeDir() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetFabricCAHomeDir")
}

// GetFabricCAID mocks base method
func (_m *MockConfig) GetFabricCAID() string {
	ret := _m.ctrl.Call(_m, "GetFabricCAID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetFabricCAID indicates an expected call of GetFabricCAID
func (_mr *MockConfigMockRecorder) GetFabricCAID() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetFabricCAID")
}

// GetFabricCAMspDir mocks base method
func (_m *MockConfig) GetFabricCAMspDir() string {
	ret := _m.ctrl.Call(_m, "GetFabricCAMspDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetFabricCAMspDir indicates an expected call of GetFabricCAMspDir
func (_mr *MockConfigMockRecorder) GetFabricCAMspDir() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetFabricCAMspDir")
}

// GetFabricCAName mocks base method
func (_m *MockConfig) GetFabricCAName() string {
	ret := _m.ctrl.Call(_m, "GetFabricCAName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetFabricCAName indicates an expected call of GetFabricCAName
func (_mr *MockConfigMockRecorder) GetFabricCAName() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetFabricCAName")
}

// GetFabricCATLSEnabledFlag mocks base method
func (_m *MockConfig) GetFabricCATLSEnabledFlag() bool {
	ret := _m.ctrl.Call(_m, "GetFabricCATLSEnabledFlag")
	ret0, _ := ret[0].(bool)
	return ret0
}

// GetFabricCATLSEnabledFlag indicates an expected call of GetFabricCATLSEnabledFlag
func (_mr *MockConfigMockRecorder) GetFabricCATLSEnabledFlag() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetFabricCATLSEnabledFlag")
}

// GetFabricClientViper mocks base method
func (_m *MockConfig) GetFabricClientViper() *viper.Viper {
	ret := _m.ctrl.Call(_m, "GetFabricClientViper")
	ret0, _ := ret[0].(*viper.Viper)
	return ret0
}

// GetFabricClientViper indicates an expected call of GetFabricClientViper
func (_mr *MockConfigMockRecorder) GetFabricClientViper() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetFabricClientViper")
}

// GetKeyStorePath mocks base method
func (_m *MockConfig) GetKeyStorePath() string {
	ret := _m.ctrl.Call(_m, "GetKeyStorePath")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetKeyStorePath indicates an expected call of GetKeyStorePath
func (_mr *MockConfigMockRecorder) GetKeyStorePath() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetKeyStorePath")
}

// GetOrdererHost mocks base method
func (_m *MockConfig) GetOrdererHost() string {
	ret := _m.ctrl.Call(_m, "GetOrdererHost")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetOrdererHost indicates an expected call of GetOrdererHost
func (_mr *MockConfigMockRecorder) GetOrdererHost() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetOrdererHost")
}

// GetOrdererPort mocks base method
func (_m *MockConfig) GetOrdererPort() string {
	ret := _m.ctrl.Call(_m, "GetOrdererPort")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetOrdererPort indicates an expected call of GetOrdererPort
func (_mr *MockConfigMockRecorder) GetOrdererPort() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetOrdererPort")
}

// GetOrdererTLSCertificate mocks base method
func (_m *MockConfig) GetOrdererTLSCertificate() string {
	ret := _m.ctrl.Call(_m, "GetOrdererTLSCertificate")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetOrdererTLSCertificate indicates an expected call of GetOrdererTLSCertificate
func (_mr *MockConfigMockRecorder) GetOrdererTLSCertificate() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetOrdererTLSCertificate")
}

// GetOrdererTLSServerHostOverride mocks base method
func (_m *MockConfig) GetOrdererTLSServerHostOverride() string {
	ret := _m.ctrl.Call(_m, "GetOrdererTLSServerHostOverride")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetOrdererTLSServerHostOverride indicates an expected call of GetOrdererTLSServerHostOverride
func (_mr *MockConfigMockRecorder) GetOrdererTLSServerHostOverride() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetOrdererTLSServerHostOverride")
}

// GetPeersConfig mocks base method
func (_m *MockConfig) GetPeersConfig() ([]api.PeerConfig, error) {
	ret := _m.ctrl.Call(_m, "GetPeersConfig")
	ret0, _ := ret[0].([]api.PeerConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPeersConfig indicates an expected call of GetPeersConfig
func (_mr *MockConfigMockRecorder) GetPeersConfig() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetPeersConfig")
}

// GetSecurityAlgorithm mocks base method
func (_m *MockConfig) GetSecurityAlgorithm() string {
	ret := _m.ctrl.Call(_m, "GetSecurityAlgorithm")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSecurityAlgorithm indicates an expected call of GetSecurityAlgorithm
func (_mr *MockConfigMockRecorder) GetSecurityAlgorithm() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSecurityAlgorithm")
}

// GetSecurityLevel mocks base method
func (_m *MockConfig) GetSecurityLevel() int {
	ret := _m.ctrl.Call(_m, "GetSecurityLevel")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetSecurityLevel indicates an expected call of GetSecurityLevel
func (_mr *MockConfigMockRecorder) GetSecurityLevel() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSecurityLevel")
}

// GetServerCertFiles mocks base method
func (_m *MockConfig) GetServerCertFiles() []string {
	ret := _m.ctrl.Call(_m, "GetServerCertFiles")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetServerCertFiles indicates an expected call of GetServerCertFiles
func (_mr *MockConfigMockRecorder) GetServerCertFiles() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetServerCertFiles")
}

// GetServerURL mocks base method
func (_m *MockConfig) GetServerURL() string {
	ret := _m.ctrl.Call(_m, "GetServerURL")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetServerURL indicates an expected call of GetServerURL
func (_mr *MockConfigMockRecorder) GetServerURL() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetServerURL")
}

// GetTLSCACertPool mocks base method
func (_m *MockConfig) GetTLSCACertPool(_param0 string) (*x509.CertPool, error) {
	ret := _m.ctrl.Call(_m, "GetTLSCACertPool", _param0)
	ret0, _ := ret[0].(*x509.CertPool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTLSCACertPool indicates an expected call of GetTLSCACertPool
func (_mr *MockConfigMockRecorder) GetTLSCACertPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTLSCACertPool", arg0)
}

// GetTLSCACertPoolFromRoots mocks base method
func (_m *MockConfig) GetTLSCACertPoolFromRoots(_param0 [][]byte) (*x509.CertPool, error) {
	ret := _m.ctrl.Call(_m, "GetTLSCACertPoolFromRoots", _param0)
	ret0, _ := ret[0].(*x509.CertPool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTLSCACertPoolFromRoots indicates an expected call of GetTLSCACertPoolFromRoots
func (_mr *MockConfigMockRecorder) GetTLSCACertPoolFromRoots(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTLSCACertPoolFromRoots", arg0)
}

// IsSecurityEnabled mocks base method
func (_m *MockConfig) IsSecurityEnabled() bool {
	ret := _m.ctrl.Call(_m, "IsSecurityEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSecurityEnabled indicates an expected call of IsSecurityEnabled
func (_mr *MockConfigMockRecorder) IsSecurityEnabled() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsSecurityEnabled")
}

// IsTLSEnabled mocks base method
func (_m *MockConfig) IsTLSEnabled() bool {
	ret := _m.ctrl.Call(_m, "IsTLSEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsTLSEnabled indicates an expected call of IsTLSEnabled
func (_mr *MockConfigMockRecorder) IsTLSEnabled() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsTLSEnabled")
}

// TcertBatchSize mocks base method
func (_m *MockConfig) TcertBatchSize() int {
	ret := _m.ctrl.Call(_m, "TcertBatchSize")
	ret0, _ := ret[0].(int)
	return ret0
}

// TcertBatchSize indicates an expected call of TcertBatchSize
func (_mr *MockConfigMockRecorder) TcertBatchSize() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TcertBatchSize")
}

// Code generated by MockGen. DO NOT EDIT.
// Source: metadata_handler.go

// Package mock_controllers is a generated GoMock package.
package mock_controllers

import (
	gomock "github.com/golang/mock/gomock"
	go_uuid "github.com/satori/go.uuid"
	metadatapb "pixielabs.ai/pixielabs/src/shared/k8s/metadatapb"
	types "pixielabs.ai/pixielabs/src/shared/types"
	controllers "pixielabs.ai/pixielabs/src/vizier/services/metadata/controllers"
	agentpb "pixielabs.ai/pixielabs/src/vizier/services/shared/agentpb"
	reflect "reflect"
)

// MockMetadataStore is a mock of MetadataStore interface
type MockMetadataStore struct {
	ctrl     *gomock.Controller
	recorder *MockMetadataStoreMockRecorder
}

// MockMetadataStoreMockRecorder is the mock recorder for MockMetadataStore
type MockMetadataStoreMockRecorder struct {
	mock *MockMetadataStore
}

// NewMockMetadataStore creates a new mock instance
func NewMockMetadataStore(ctrl *gomock.Controller) *MockMetadataStore {
	mock := &MockMetadataStore{ctrl: ctrl}
	mock.recorder = &MockMetadataStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMetadataStore) EXPECT() *MockMetadataStoreMockRecorder {
	return m.recorder
}

// GetClusterCIDR mocks base method
func (m *MockMetadataStore) GetClusterCIDR() string {
	ret := m.ctrl.Call(m, "GetClusterCIDR")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetClusterCIDR indicates an expected call of GetClusterCIDR
func (mr *MockMetadataStoreMockRecorder) GetClusterCIDR() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterCIDR", reflect.TypeOf((*MockMetadataStore)(nil).GetClusterCIDR))
}

// GetServiceCIDR mocks base method
func (m *MockMetadataStore) GetServiceCIDR() string {
	ret := m.ctrl.Call(m, "GetServiceCIDR")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetServiceCIDR indicates an expected call of GetServiceCIDR
func (mr *MockMetadataStoreMockRecorder) GetServiceCIDR() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceCIDR", reflect.TypeOf((*MockMetadataStore)(nil).GetServiceCIDR))
}

// GetAgent mocks base method
func (m *MockMetadataStore) GetAgent(agentID go_uuid.UUID) (*agentpb.Agent, error) {
	ret := m.ctrl.Call(m, "GetAgent", agentID)
	ret0, _ := ret[0].(*agentpb.Agent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAgent indicates an expected call of GetAgent
func (mr *MockMetadataStoreMockRecorder) GetAgent(agentID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgent", reflect.TypeOf((*MockMetadataStore)(nil).GetAgent), agentID)
}

// GetAgentIDForHostname mocks base method
func (m *MockMetadataStore) GetAgentIDForHostname(hostname string) (string, error) {
	ret := m.ctrl.Call(m, "GetAgentIDForHostname", hostname)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAgentIDForHostname indicates an expected call of GetAgentIDForHostname
func (mr *MockMetadataStoreMockRecorder) GetAgentIDForHostname(hostname interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgentIDForHostname", reflect.TypeOf((*MockMetadataStore)(nil).GetAgentIDForHostname), hostname)
}

// DeleteAgent mocks base method
func (m *MockMetadataStore) DeleteAgent(agentID go_uuid.UUID) error {
	ret := m.ctrl.Call(m, "DeleteAgent", agentID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAgent indicates an expected call of DeleteAgent
func (mr *MockMetadataStoreMockRecorder) DeleteAgent(agentID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAgent", reflect.TypeOf((*MockMetadataStore)(nil).DeleteAgent), agentID)
}

// CreateAgent mocks base method
func (m *MockMetadataStore) CreateAgent(agentID go_uuid.UUID, a *agentpb.Agent) error {
	ret := m.ctrl.Call(m, "CreateAgent", agentID, a)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAgent indicates an expected call of CreateAgent
func (mr *MockMetadataStoreMockRecorder) CreateAgent(agentID, a interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAgent", reflect.TypeOf((*MockMetadataStore)(nil).CreateAgent), agentID, a)
}

// UpdateAgent mocks base method
func (m *MockMetadataStore) UpdateAgent(agentID go_uuid.UUID, a *agentpb.Agent) error {
	ret := m.ctrl.Call(m, "UpdateAgent", agentID, a)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAgent indicates an expected call of UpdateAgent
func (mr *MockMetadataStoreMockRecorder) UpdateAgent(agentID, a interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAgent", reflect.TypeOf((*MockMetadataStore)(nil).UpdateAgent), agentID, a)
}

// GetAgents mocks base method
func (m *MockMetadataStore) GetAgents() ([]*agentpb.Agent, error) {
	ret := m.ctrl.Call(m, "GetAgents")
	ret0, _ := ret[0].([]*agentpb.Agent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAgents indicates an expected call of GetAgents
func (mr *MockMetadataStoreMockRecorder) GetAgents() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgents", reflect.TypeOf((*MockMetadataStore)(nil).GetAgents))
}

// GetAgentsForHostnames mocks base method
func (m *MockMetadataStore) GetAgentsForHostnames(arg0 *[]string) ([]string, error) {
	ret := m.ctrl.Call(m, "GetAgentsForHostnames", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAgentsForHostnames indicates an expected call of GetAgentsForHostnames
func (mr *MockMetadataStoreMockRecorder) GetAgentsForHostnames(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgentsForHostnames", reflect.TypeOf((*MockMetadataStore)(nil).GetAgentsForHostnames), arg0)
}

// GetASID mocks base method
func (m *MockMetadataStore) GetASID() (uint32, error) {
	ret := m.ctrl.Call(m, "GetASID")
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetASID indicates an expected call of GetASID
func (mr *MockMetadataStoreMockRecorder) GetASID() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetASID", reflect.TypeOf((*MockMetadataStore)(nil).GetASID))
}

// GetKelvinIDs mocks base method
func (m *MockMetadataStore) GetKelvinIDs() ([]string, error) {
	ret := m.ctrl.Call(m, "GetKelvinIDs")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKelvinIDs indicates an expected call of GetKelvinIDs
func (mr *MockMetadataStoreMockRecorder) GetKelvinIDs() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKelvinIDs", reflect.TypeOf((*MockMetadataStore)(nil).GetKelvinIDs))
}

// GetComputedSchemas mocks base method
func (m *MockMetadataStore) GetComputedSchemas() ([]*metadatapb.SchemaInfo, error) {
	ret := m.ctrl.Call(m, "GetComputedSchemas")
	ret0, _ := ret[0].([]*metadatapb.SchemaInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComputedSchemas indicates an expected call of GetComputedSchemas
func (mr *MockMetadataStoreMockRecorder) GetComputedSchemas() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComputedSchemas", reflect.TypeOf((*MockMetadataStore)(nil).GetComputedSchemas))
}

// UpdateSchemas mocks base method
func (m *MockMetadataStore) UpdateSchemas(agentID go_uuid.UUID, schemas []*metadatapb.SchemaInfo) error {
	ret := m.ctrl.Call(m, "UpdateSchemas", agentID, schemas)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSchemas indicates an expected call of UpdateSchemas
func (mr *MockMetadataStoreMockRecorder) UpdateSchemas(agentID, schemas interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSchemas", reflect.TypeOf((*MockMetadataStore)(nil).UpdateSchemas), agentID, schemas)
}

// UpdateProcesses mocks base method
func (m *MockMetadataStore) UpdateProcesses(processes []*metadatapb.ProcessInfo) error {
	ret := m.ctrl.Call(m, "UpdateProcesses", processes)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProcesses indicates an expected call of UpdateProcesses
func (mr *MockMetadataStoreMockRecorder) UpdateProcesses(processes interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProcesses", reflect.TypeOf((*MockMetadataStore)(nil).UpdateProcesses), processes)
}

// GetProcesses mocks base method
func (m *MockMetadataStore) GetProcesses(upids []*types.UInt128) ([]*metadatapb.ProcessInfo, error) {
	ret := m.ctrl.Call(m, "GetProcesses", upids)
	ret0, _ := ret[0].([]*metadatapb.ProcessInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProcesses indicates an expected call of GetProcesses
func (mr *MockMetadataStoreMockRecorder) GetProcesses(upids interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProcesses", reflect.TypeOf((*MockMetadataStore)(nil).GetProcesses), upids)
}

// GetPods mocks base method
func (m *MockMetadataStore) GetPods() ([]*metadatapb.Pod, error) {
	ret := m.ctrl.Call(m, "GetPods")
	ret0, _ := ret[0].([]*metadatapb.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPods indicates an expected call of GetPods
func (mr *MockMetadataStoreMockRecorder) GetPods() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPods", reflect.TypeOf((*MockMetadataStore)(nil).GetPods))
}

// UpdatePod mocks base method
func (m *MockMetadataStore) UpdatePod(arg0 *metadatapb.Pod, arg1 bool) error {
	ret := m.ctrl.Call(m, "UpdatePod", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePod indicates an expected call of UpdatePod
func (mr *MockMetadataStoreMockRecorder) UpdatePod(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePod", reflect.TypeOf((*MockMetadataStore)(nil).UpdatePod), arg0, arg1)
}

// GetNodePods mocks base method
func (m *MockMetadataStore) GetNodePods(hostname string) ([]*metadatapb.Pod, error) {
	ret := m.ctrl.Call(m, "GetNodePods", hostname)
	ret0, _ := ret[0].([]*metadatapb.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNodePods indicates an expected call of GetNodePods
func (mr *MockMetadataStoreMockRecorder) GetNodePods(hostname interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodePods", reflect.TypeOf((*MockMetadataStore)(nil).GetNodePods), hostname)
}

// GetEndpoints mocks base method
func (m *MockMetadataStore) GetEndpoints() ([]*metadatapb.Endpoints, error) {
	ret := m.ctrl.Call(m, "GetEndpoints")
	ret0, _ := ret[0].([]*metadatapb.Endpoints)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEndpoints indicates an expected call of GetEndpoints
func (mr *MockMetadataStoreMockRecorder) GetEndpoints() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEndpoints", reflect.TypeOf((*MockMetadataStore)(nil).GetEndpoints))
}

// UpdateEndpoints mocks base method
func (m *MockMetadataStore) UpdateEndpoints(arg0 *metadatapb.Endpoints, arg1 bool) error {
	ret := m.ctrl.Call(m, "UpdateEndpoints", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEndpoints indicates an expected call of UpdateEndpoints
func (mr *MockMetadataStoreMockRecorder) UpdateEndpoints(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEndpoints", reflect.TypeOf((*MockMetadataStore)(nil).UpdateEndpoints), arg0, arg1)
}

// GetNodeEndpoints mocks base method
func (m *MockMetadataStore) GetNodeEndpoints(hostname string) ([]*metadatapb.Endpoints, error) {
	ret := m.ctrl.Call(m, "GetNodeEndpoints", hostname)
	ret0, _ := ret[0].([]*metadatapb.Endpoints)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNodeEndpoints indicates an expected call of GetNodeEndpoints
func (mr *MockMetadataStoreMockRecorder) GetNodeEndpoints(hostname interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodeEndpoints", reflect.TypeOf((*MockMetadataStore)(nil).GetNodeEndpoints), hostname)
}

// GetServices mocks base method
func (m *MockMetadataStore) GetServices() ([]*metadatapb.Service, error) {
	ret := m.ctrl.Call(m, "GetServices")
	ret0, _ := ret[0].([]*metadatapb.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServices indicates an expected call of GetServices
func (mr *MockMetadataStoreMockRecorder) GetServices() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServices", reflect.TypeOf((*MockMetadataStore)(nil).GetServices))
}

// UpdateService mocks base method
func (m *MockMetadataStore) UpdateService(arg0 *metadatapb.Service, arg1 bool) error {
	ret := m.ctrl.Call(m, "UpdateService", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateService indicates an expected call of UpdateService
func (mr *MockMetadataStoreMockRecorder) UpdateService(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateService", reflect.TypeOf((*MockMetadataStore)(nil).UpdateService), arg0, arg1)
}

// GetContainers mocks base method
func (m *MockMetadataStore) GetContainers() ([]*metadatapb.ContainerInfo, error) {
	ret := m.ctrl.Call(m, "GetContainers")
	ret0, _ := ret[0].([]*metadatapb.ContainerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContainers indicates an expected call of GetContainers
func (mr *MockMetadataStoreMockRecorder) GetContainers() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainers", reflect.TypeOf((*MockMetadataStore)(nil).GetContainers))
}

// UpdateContainer mocks base method
func (m *MockMetadataStore) UpdateContainer(arg0 *metadatapb.ContainerInfo) error {
	ret := m.ctrl.Call(m, "UpdateContainer", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateContainer indicates an expected call of UpdateContainer
func (mr *MockMetadataStoreMockRecorder) UpdateContainer(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateContainer", reflect.TypeOf((*MockMetadataStore)(nil).UpdateContainer), arg0)
}

// UpdateContainersFromPod mocks base method
func (m *MockMetadataStore) UpdateContainersFromPod(arg0 *metadatapb.Pod, arg1 bool) error {
	ret := m.ctrl.Call(m, "UpdateContainersFromPod", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateContainersFromPod indicates an expected call of UpdateContainersFromPod
func (mr *MockMetadataStoreMockRecorder) UpdateContainersFromPod(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateContainersFromPod", reflect.TypeOf((*MockMetadataStore)(nil).UpdateContainersFromPod), arg0, arg1)
}

// GetMetadataUpdates mocks base method
func (m *MockMetadataStore) GetMetadataUpdates(hostname string) ([]*metadatapb.ResourceUpdate, error) {
	ret := m.ctrl.Call(m, "GetMetadataUpdates", hostname)
	ret0, _ := ret[0].([]*metadatapb.ResourceUpdate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetadataUpdates indicates an expected call of GetMetadataUpdates
func (mr *MockMetadataStoreMockRecorder) GetMetadataUpdates(hostname interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetadataUpdates", reflect.TypeOf((*MockMetadataStore)(nil).GetMetadataUpdates), hostname)
}

// MockMetadataSubscriber is a mock of MetadataSubscriber interface
type MockMetadataSubscriber struct {
	ctrl     *gomock.Controller
	recorder *MockMetadataSubscriberMockRecorder
}

// MockMetadataSubscriberMockRecorder is the mock recorder for MockMetadataSubscriber
type MockMetadataSubscriberMockRecorder struct {
	mock *MockMetadataSubscriber
}

// NewMockMetadataSubscriber creates a new mock instance
func NewMockMetadataSubscriber(ctrl *gomock.Controller) *MockMetadataSubscriber {
	mock := &MockMetadataSubscriber{ctrl: ctrl}
	mock.recorder = &MockMetadataSubscriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMetadataSubscriber) EXPECT() *MockMetadataSubscriberMockRecorder {
	return m.recorder
}

// HandleUpdate mocks base method
func (m *MockMetadataSubscriber) HandleUpdate(arg0 *controllers.UpdateMessage) {
	m.ctrl.Call(m, "HandleUpdate", arg0)
}

// HandleUpdate indicates an expected call of HandleUpdate
func (mr *MockMetadataSubscriberMockRecorder) HandleUpdate(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleUpdate", reflect.TypeOf((*MockMetadataSubscriber)(nil).HandleUpdate), arg0)
}

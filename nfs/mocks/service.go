// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dell/csm-sharednfs/nfs (interfaces: Service)
//
// Generated by this command:
//
//	mockgen -destination=mocks/service.go -package=mocks github.com/dell/csm-sharednfs/nfs Service
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	net "net"
	reflect "reflect"

	csi "github.com/container-storage-interface/spec/lib/go/csi"
	gocsi "github.com/dell/gocsi"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
	isgomock struct{}
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// BeforeServe mocks base method.
func (m *MockService) BeforeServe(arg0 context.Context, arg1 *gocsi.StoragePlugin, arg2 net.Listener) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeforeServe", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// BeforeServe indicates an expected call of BeforeServe.
func (mr *MockServiceMockRecorder) BeforeServe(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeServe", reflect.TypeOf((*MockService)(nil).BeforeServe), arg0, arg1, arg2)
}

// ControllerExpandVolume mocks base method.
func (m *MockService) ControllerExpandVolume(arg0 context.Context, arg1 *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerExpandVolume", arg0, arg1)
	ret0, _ := ret[0].(*csi.ControllerExpandVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerExpandVolume indicates an expected call of ControllerExpandVolume.
func (mr *MockServiceMockRecorder) ControllerExpandVolume(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerExpandVolume", reflect.TypeOf((*MockService)(nil).ControllerExpandVolume), arg0, arg1)
}

// ControllerGetCapabilities mocks base method.
func (m *MockService) ControllerGetCapabilities(arg0 context.Context, arg1 *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerGetCapabilities", arg0, arg1)
	ret0, _ := ret[0].(*csi.ControllerGetCapabilitiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerGetCapabilities indicates an expected call of ControllerGetCapabilities.
func (mr *MockServiceMockRecorder) ControllerGetCapabilities(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerGetCapabilities", reflect.TypeOf((*MockService)(nil).ControllerGetCapabilities), arg0, arg1)
}

// ControllerGetVolume mocks base method.
func (m *MockService) ControllerGetVolume(arg0 context.Context, arg1 *csi.ControllerGetVolumeRequest) (*csi.ControllerGetVolumeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerGetVolume", arg0, arg1)
	ret0, _ := ret[0].(*csi.ControllerGetVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerGetVolume indicates an expected call of ControllerGetVolume.
func (mr *MockServiceMockRecorder) ControllerGetVolume(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerGetVolume", reflect.TypeOf((*MockService)(nil).ControllerGetVolume), arg0, arg1)
}

// ControllerPublishVolume mocks base method.
func (m *MockService) ControllerPublishVolume(arg0 context.Context, arg1 *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerPublishVolume", arg0, arg1)
	ret0, _ := ret[0].(*csi.ControllerPublishVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerPublishVolume indicates an expected call of ControllerPublishVolume.
func (mr *MockServiceMockRecorder) ControllerPublishVolume(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerPublishVolume", reflect.TypeOf((*MockService)(nil).ControllerPublishVolume), arg0, arg1)
}

// ControllerUnpublishVolume mocks base method.
func (m *MockService) ControllerUnpublishVolume(arg0 context.Context, arg1 *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerUnpublishVolume", arg0, arg1)
	ret0, _ := ret[0].(*csi.ControllerUnpublishVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerUnpublishVolume indicates an expected call of ControllerUnpublishVolume.
func (mr *MockServiceMockRecorder) ControllerUnpublishVolume(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerUnpublishVolume", reflect.TypeOf((*MockService)(nil).ControllerUnpublishVolume), arg0, arg1)
}

// CreateSnapshot mocks base method.
func (m *MockService) CreateSnapshot(arg0 context.Context, arg1 *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSnapshot", arg0, arg1)
	ret0, _ := ret[0].(*csi.CreateSnapshotResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSnapshot indicates an expected call of CreateSnapshot.
func (mr *MockServiceMockRecorder) CreateSnapshot(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSnapshot", reflect.TypeOf((*MockService)(nil).CreateSnapshot), arg0, arg1)
}

// CreateVolume mocks base method.
func (m *MockService) CreateVolume(arg0 context.Context, arg1 *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVolume", arg0, arg1)
	ret0, _ := ret[0].(*csi.CreateVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVolume indicates an expected call of CreateVolume.
func (mr *MockServiceMockRecorder) CreateVolume(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVolume", reflect.TypeOf((*MockService)(nil).CreateVolume), arg0, arg1)
}

// DeleteSnapshot mocks base method.
func (m *MockService) DeleteSnapshot(arg0 context.Context, arg1 *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSnapshot", arg0, arg1)
	ret0, _ := ret[0].(*csi.DeleteSnapshotResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSnapshot indicates an expected call of DeleteSnapshot.
func (mr *MockServiceMockRecorder) DeleteSnapshot(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSnapshot", reflect.TypeOf((*MockService)(nil).DeleteSnapshot), arg0, arg1)
}

// DeleteVolume mocks base method.
func (m *MockService) DeleteVolume(arg0 context.Context, arg1 *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVolume", arg0, arg1)
	ret0, _ := ret[0].(*csi.DeleteVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteVolume indicates an expected call of DeleteVolume.
func (mr *MockServiceMockRecorder) DeleteVolume(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVolume", reflect.TypeOf((*MockService)(nil).DeleteVolume), arg0, arg1)
}

// GetCapacity mocks base method.
func (m *MockService) GetCapacity(arg0 context.Context, arg1 *csi.GetCapacityRequest) (*csi.GetCapacityResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCapacity", arg0, arg1)
	ret0, _ := ret[0].(*csi.GetCapacityResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCapacity indicates an expected call of GetCapacity.
func (mr *MockServiceMockRecorder) GetCapacity(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCapacity", reflect.TypeOf((*MockService)(nil).GetCapacity), arg0, arg1)
}

// GetPluginCapabilities mocks base method.
func (m *MockService) GetPluginCapabilities(arg0 context.Context, arg1 *csi.GetPluginCapabilitiesRequest) (*csi.GetPluginCapabilitiesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPluginCapabilities", arg0, arg1)
	ret0, _ := ret[0].(*csi.GetPluginCapabilitiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPluginCapabilities indicates an expected call of GetPluginCapabilities.
func (mr *MockServiceMockRecorder) GetPluginCapabilities(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPluginCapabilities", reflect.TypeOf((*MockService)(nil).GetPluginCapabilities), arg0, arg1)
}

// GetPluginInfo mocks base method.
func (m *MockService) GetPluginInfo(arg0 context.Context, arg1 *csi.GetPluginInfoRequest) (*csi.GetPluginInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPluginInfo", arg0, arg1)
	ret0, _ := ret[0].(*csi.GetPluginInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPluginInfo indicates an expected call of GetPluginInfo.
func (mr *MockServiceMockRecorder) GetPluginInfo(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPluginInfo", reflect.TypeOf((*MockService)(nil).GetPluginInfo), arg0, arg1)
}

// ListSnapshots mocks base method.
func (m *MockService) ListSnapshots(arg0 context.Context, arg1 *csi.ListSnapshotsRequest) (*csi.ListSnapshotsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSnapshots", arg0, arg1)
	ret0, _ := ret[0].(*csi.ListSnapshotsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSnapshots indicates an expected call of ListSnapshots.
func (mr *MockServiceMockRecorder) ListSnapshots(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSnapshots", reflect.TypeOf((*MockService)(nil).ListSnapshots), arg0, arg1)
}

// ListVolumes mocks base method.
func (m *MockService) ListVolumes(arg0 context.Context, arg1 *csi.ListVolumesRequest) (*csi.ListVolumesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVolumes", arg0, arg1)
	ret0, _ := ret[0].(*csi.ListVolumesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVolumes indicates an expected call of ListVolumes.
func (mr *MockServiceMockRecorder) ListVolumes(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVolumes", reflect.TypeOf((*MockService)(nil).ListVolumes), arg0, arg1)
}

// MountVolume mocks base method.
func (m *MockService) MountVolume(arg0 context.Context, arg1, arg2, arg3 string, arg4 map[string]string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MountVolume", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MountVolume indicates an expected call of MountVolume.
func (mr *MockServiceMockRecorder) MountVolume(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MountVolume", reflect.TypeOf((*MockService)(nil).MountVolume), arg0, arg1, arg2, arg3, arg4)
}

// NodeExpandVolume mocks base method.
func (m *MockService) NodeExpandVolume(arg0 context.Context, arg1 *csi.NodeExpandVolumeRequest) (*csi.NodeExpandVolumeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeExpandVolume", arg0, arg1)
	ret0, _ := ret[0].(*csi.NodeExpandVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NodeExpandVolume indicates an expected call of NodeExpandVolume.
func (mr *MockServiceMockRecorder) NodeExpandVolume(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeExpandVolume", reflect.TypeOf((*MockService)(nil).NodeExpandVolume), arg0, arg1)
}

// NodeGetCapabilities mocks base method.
func (m *MockService) NodeGetCapabilities(arg0 context.Context, arg1 *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeGetCapabilities", arg0, arg1)
	ret0, _ := ret[0].(*csi.NodeGetCapabilitiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NodeGetCapabilities indicates an expected call of NodeGetCapabilities.
func (mr *MockServiceMockRecorder) NodeGetCapabilities(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeGetCapabilities", reflect.TypeOf((*MockService)(nil).NodeGetCapabilities), arg0, arg1)
}

// NodeGetInfo mocks base method.
func (m *MockService) NodeGetInfo(arg0 context.Context, arg1 *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeGetInfo", arg0, arg1)
	ret0, _ := ret[0].(*csi.NodeGetInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NodeGetInfo indicates an expected call of NodeGetInfo.
func (mr *MockServiceMockRecorder) NodeGetInfo(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeGetInfo", reflect.TypeOf((*MockService)(nil).NodeGetInfo), arg0, arg1)
}

// NodeGetVolumeStats mocks base method.
func (m *MockService) NodeGetVolumeStats(arg0 context.Context, arg1 *csi.NodeGetVolumeStatsRequest) (*csi.NodeGetVolumeStatsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeGetVolumeStats", arg0, arg1)
	ret0, _ := ret[0].(*csi.NodeGetVolumeStatsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NodeGetVolumeStats indicates an expected call of NodeGetVolumeStats.
func (mr *MockServiceMockRecorder) NodeGetVolumeStats(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeGetVolumeStats", reflect.TypeOf((*MockService)(nil).NodeGetVolumeStats), arg0, arg1)
}

// NodePublishVolume mocks base method.
func (m *MockService) NodePublishVolume(arg0 context.Context, arg1 *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodePublishVolume", arg0, arg1)
	ret0, _ := ret[0].(*csi.NodePublishVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NodePublishVolume indicates an expected call of NodePublishVolume.
func (mr *MockServiceMockRecorder) NodePublishVolume(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodePublishVolume", reflect.TypeOf((*MockService)(nil).NodePublishVolume), arg0, arg1)
}

// NodeStageVolume mocks base method.
func (m *MockService) NodeStageVolume(arg0 context.Context, arg1 *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeStageVolume", arg0, arg1)
	ret0, _ := ret[0].(*csi.NodeStageVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NodeStageVolume indicates an expected call of NodeStageVolume.
func (mr *MockServiceMockRecorder) NodeStageVolume(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeStageVolume", reflect.TypeOf((*MockService)(nil).NodeStageVolume), arg0, arg1)
}

// NodeUnpublishVolume mocks base method.
func (m *MockService) NodeUnpublishVolume(arg0 context.Context, arg1 *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeUnpublishVolume", arg0, arg1)
	ret0, _ := ret[0].(*csi.NodeUnpublishVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NodeUnpublishVolume indicates an expected call of NodeUnpublishVolume.
func (mr *MockServiceMockRecorder) NodeUnpublishVolume(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeUnpublishVolume", reflect.TypeOf((*MockService)(nil).NodeUnpublishVolume), arg0, arg1)
}

// NodeUnstageVolume mocks base method.
func (m *MockService) NodeUnstageVolume(arg0 context.Context, arg1 *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeUnstageVolume", arg0, arg1)
	ret0, _ := ret[0].(*csi.NodeUnstageVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NodeUnstageVolume indicates an expected call of NodeUnstageVolume.
func (mr *MockServiceMockRecorder) NodeUnstageVolume(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeUnstageVolume", reflect.TypeOf((*MockService)(nil).NodeUnstageVolume), arg0, arg1)
}

// Probe mocks base method.
func (m *MockService) Probe(arg0 context.Context, arg1 *csi.ProbeRequest) (*csi.ProbeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Probe", arg0, arg1)
	ret0, _ := ret[0].(*csi.ProbeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Probe indicates an expected call of Probe.
func (mr *MockServiceMockRecorder) Probe(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Probe", reflect.TypeOf((*MockService)(nil).Probe), arg0, arg1)
}

// ProcessMapSecretChange mocks base method.
func (m *MockService) ProcessMapSecretChange() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessMapSecretChange")
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessMapSecretChange indicates an expected call of ProcessMapSecretChange.
func (mr *MockServiceMockRecorder) ProcessMapSecretChange() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessMapSecretChange", reflect.TypeOf((*MockService)(nil).ProcessMapSecretChange))
}

// RegisterAdditionalServers mocks base method.
func (m *MockService) RegisterAdditionalServers(server *grpc.Server) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterAdditionalServers", server)
}

// RegisterAdditionalServers indicates an expected call of RegisterAdditionalServers.
func (mr *MockServiceMockRecorder) RegisterAdditionalServers(server any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterAdditionalServers", reflect.TypeOf((*MockService)(nil).RegisterAdditionalServers), server)
}

// UnmountVolume mocks base method.
func (m *MockService) UnmountVolume(arg0 context.Context, arg1, arg2 string, arg3 map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnmountVolume", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnmountVolume indicates an expected call of UnmountVolume.
func (mr *MockServiceMockRecorder) UnmountVolume(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnmountVolume", reflect.TypeOf((*MockService)(nil).UnmountVolume), arg0, arg1, arg2, arg3)
}

// ValidateVolumeCapabilities mocks base method.
func (m *MockService) ValidateVolumeCapabilities(arg0 context.Context, arg1 *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateVolumeCapabilities", arg0, arg1)
	ret0, _ := ret[0].(*csi.ValidateVolumeCapabilitiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateVolumeCapabilities indicates an expected call of ValidateVolumeCapabilities.
func (mr *MockServiceMockRecorder) ValidateVolumeCapabilities(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateVolumeCapabilities", reflect.TypeOf((*MockService)(nil).ValidateVolumeCapabilities), arg0, arg1)
}

// VolumeIDToArrayID mocks base method.
func (m *MockService) VolumeIDToArrayID(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VolumeIDToArrayID", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// VolumeIDToArrayID indicates an expected call of VolumeIDToArrayID.
func (mr *MockServiceMockRecorder) VolumeIDToArrayID(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeIDToArrayID", reflect.TypeOf((*MockService)(nil).VolumeIDToArrayID), arg0)
}

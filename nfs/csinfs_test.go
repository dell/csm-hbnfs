/*
Copyright © 2025 Dell Inc. or its subsidiaries. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package nfs

import (
	"context"
	"errors"
	"net"
	"reflect"
	"strings"
	"testing"

	csictx "github.com/dell/gocsi/context"

	k8s "github.com/dell/csm-hbnfs/nfs/k8s"
	"github.com/dell/csm-hbnfs/nfs/mocks"
	"github.com/dell/gocsi"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/metadata"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/rest"
)

func TestIsNFSStorageClass(t *testing.T) {
	type args struct {
		parameters map[string]string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "is nfs storage class",
			args: args{
				parameters: map[string]string{
					"csi-nfs": "RWX",
				},
			},
			want: true,
		},
		{
			name: "is not nfs storage class",
			args: args{
				parameters: map[string]string{
					"csi-nfs": "RWO",
				},
			},
			want: false,
		},
		{
			name: "parameters missing key 'csi-nfs' key",
			args: args{
				parameters: map[string]string{
					"missing-key": "RWX",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNFSStorageClass(tt.args.parameters); got != tt.want {
				t.Errorf("IsNFSStorageClass() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasNFSPrefix(t *testing.T) {
	type args struct {
		volumeID string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "has nfs prefix",
			args: args{
				volumeID: "nfs-abcde01234",
			},
			want: true,
		},
		{
			name: "does not have nfs prefix",
			args: args{
				volumeID: "abcde01234",
			},
			want: false,
		},
		{
			name: "contains nfs prefix but not as a prefix",
			args: args{
				volumeID: "abcde-nfs-01234",
			},
			want: false,
		},
		{
			name: "empty string",
			args: args{
				volumeID: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasNFSPrefix(tt.args.volumeID); got != tt.want {
				t.Errorf("IsNFSVolumeID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNFSToArrayVolumeID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "trims nfs prefix",
			args: args{
				id: "nfs-abcde01234",
			},
			want: "abcde01234",
		},
		{
			name: "no prefix to trim",
			args: args{
				id: "abcde01234",
			},
			want: "abcde01234",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NFSToArrayVolumeID(tt.args.id); got != tt.want {
				t.Errorf("NFSToArrayVolumeID() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestVolumeIDToServiceName(t *testing.T) {
// 	type args struct {
// 		id string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want string
// 	}{
// 		{
// 			name: "",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := VolumeIDToServiceName(tt.args.id); got != tt.want {
// 				t.Errorf("VolumeIDToServiceName() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestArrayToNFSVolumeID(t *testing.T) {
	type args struct {
		arrayid string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "adds nfs prefix",
			args: args{
				arrayid: "abcde01234",
			},
			want: "nfs-abcde01234",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayToNFSVolumeID(tt.args.arrayid); got != tt.want {
				t.Errorf("ArrayToNFSVolumeID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	var defaultSvc *CsiNfsService = nil
	afterEach := func() {
		nfsService = defaultSvc
	}

	type args struct {
		provisionerName string
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		{
			name: "create a new csi nfs service",
			args: args{
				provisionerName: "csi-powerstore.dellemc.com",
			},
			want: &CsiNfsService{
				provisionerName: "csi-powerstore.dellemc.com",
				executor:        &LocalExecutor{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer afterEach()

			if got := New(tt.args.provisionerName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPutVcsiService(t *testing.T) {
	var defaultNfsService *CsiNfsService = nil
	afterEach := func() {
		nfsService = defaultNfsService
	}

	type args struct {
		vcsi Service
	}
	tests := []struct {
		name   string
		args   args
		before func()
		want   Service
	}{
		{
			name: "save vcsi service in global nfs service",
			args: args{
				// should be the main driver's service, but for simplicity's sake, we use the csi nfs service
				vcsi: &CsiNfsService{},
			},
			before: func() {
				_ = New("csi-powerstore.dellemc.com")
			},
			want: &CsiNfsService{
				vcsi:            &CsiNfsService{},
				provisionerName: "csi-powerstore.dellemc.com",
				executor:        &LocalExecutor{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer afterEach()
			tt.before()

			PutVcsiService(tt.args.vcsi)
			if !reflect.DeepEqual(nfsService, tt.want) {
				t.Errorf("PutVcsiService() = %v, want %v", nfsService, tt.want)
			}
		})
	}
}

func TestCsiNfsService_validateGlobalVariables(t *testing.T) {
	// save and restore the defaults after each test case
	defaultNodeRoot := NodeRoot
	defaultNfsExportDir := NfsExportDirectory
	defaultDriverNamespace := DriverNamespace
	defaultDriverName := DriverName
	defaultOpSys := opSys

	afterEach := func() {
		NodeRoot = defaultNodeRoot
		NfsExportDirectory = defaultNfsExportDir
		DriverNamespace = defaultDriverNamespace
		DriverName = defaultDriverName
		opSys = defaultOpSys
	}

	type fields struct {
		mode string
	}
	tests := []struct {
		name       string
		fields     fields
		before     func()
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "mode is empty",
			fields: fields{
				mode: "",
			},
			before:     func() {},
			wantErr:    true,
			wantErrMsg: "s.mode not set to \"node\" or \"controller\"",
		},
		{
			name: "mode is node and Node Root is unset",
			fields: fields{
				mode: "node",
			},
			before: func() {
				NodeRoot = ""
			},
			wantErr:    true,
			wantErrMsg: "csi-nfs NodeRoot variable must be set; used for chroot into node; validated with /noderoot/etc/exports",
		},
		{
			name: "mode is controller and driver namespace is unset",
			fields: fields{
				mode: "controller",
			},
			before: func() {
				DriverNamespace = ""
			},
			wantErr:    true,
			wantErrMsg: "DriverNamespace variable not set; this is used to find Services and EndpointSlices in the driver namespace",
		},
		{
			name: "mode is controller and driver name is unset",
			fields: fields{
				mode: "controller",
			},
			before: func() {
				DriverNamespace = "powerstore"
				DriverName = ""
			},
			wantErr:    true,
			wantErrMsg: "DriverName not set. This is the value of the driver name in the csinode objects, e.g. csi-vxflexos.dellemc.com",
		},
		{
			name: "validate global variables in controller mode",
			fields: fields{
				mode: "controller",
			},
			before: func() {
				DriverNamespace = "powerstore"
				DriverName = "csi-powerstore.dellemc.com"
			},
			wantErr: false,
		},
		{
			name: "/etc/exports file does not exist",
			fields: fields{
				mode: "node",
			},
			before: func() {
				NodeRoot = "/noderoot"

				os := mocks.NewMockOSInterface(gomock.NewController(t))
				os.EXPECT().Stat(gomock.Any()).Times(1).Return(nil, errors.New("os stat error"))
				opSys = os
			},
			wantErr:    true,
			wantErrMsg: "could not stat NodeRoot/etc/exports - this file must exist for kernel nfs installations",
		},
		{
			name: "NfsExportsDirectory unset",
			fields: fields{
				mode: "node",
			},
			before: func() {
				NodeRoot = "/noderoot"

				os := mocks.NewMockOSInterface(gomock.NewController(t))
				os.EXPECT().Stat(gomock.Any()).Times(1).Return(nil, nil)
				opSys = os
			},
			wantErr:    true,
			wantErrMsg: "NfsExportDirectory not set it is where block devices are exported via NFS",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer afterEach()
			tt.before()

			s := &CsiNfsService{
				mode: tt.fields.mode,
			}
			err := s.validateGlobalVariables()
			if (err != nil) != tt.wantErr {
				t.Errorf("CsiNfsService.validateGlobalVariables() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr && (err != nil) {
				assert.Contains(t, err.Error(), tt.wantErrMsg)
			}
		})
	}
}

func Test_getRequestIdFromContext(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "get the request ID from the request context metadata",
			args: args{
				ctx: func() context.Context {
					return metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{"csi.requestid": "1"}))
				}(),
			},
			want: "1",
		},
		{
			name: "no request id in the context metadata",
			args: args{
				ctx: context.Background(),
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRequestIdFromContext(tt.args.ctx); got != tt.want {
				t.Errorf("getRequestIdFromContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCsiNfsService_startNodeMonitors(t *testing.T) {
	type fields struct {
		k8sclient *k8s.K8sClient
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "for controller node",
			fields: fields{
				k8sclient: func() *k8s.K8sClient {
					clientSet := fake.NewClientset()
					_, err := clientSet.CoreV1().Nodes().Create(context.Background(), &v1.Node{
						Spec: v1.NodeSpec{
							Taints: []v1.Taint{
								{
									Key: "node-role.kubernetes.io/control-plane",
								},
							},
						},
					}, metav1.CreateOptions{})
					if err != nil {
						t.Error("failed to add node via the fake client")
					}
					return &k8s.K8sClient{
						Clientset: clientSet,
					}
				}(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CsiNfsService{
				k8sclient: tt.fields.k8sclient,
			}
			if err := s.startNodeMonitors(); (err != nil) != tt.wantErr {
				t.Errorf("CsiNfsService.startNodeMonitors() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCsiNfsService_BeforeServe(t *testing.T) {
	defaultOpSys := opSys
	defaultDriverNamespace := DriverNamespace
	defaultDriverName := DriverName
	defaultNewForConfigFunc := k8s.NewForConfigFunc
	defaultRestInClusterConfigFunc := k8s.RestInClusterConfigFunc
	// restore defaults after each tests case
	afterEach := func() {
		opSys = defaultOpSys
		DriverNamespace = defaultDriverNamespace
		DriverName = defaultDriverName
		k8s.NewForConfigFunc = defaultNewForConfigFunc
		k8s.RestInClusterConfigFunc = defaultRestInClusterConfigFunc
	}

	type fields struct {
		vcsi            Service
		md              Service
		provisionerName string
		mode            string
		nodeID          string
		nodeIPAddress   string
		podCIDR         string
		nodeName        string
		k8sclient       *k8s.K8sClient
		executor        Executor
	}
	type args struct {
		ctx context.Context
		sp  *gocsi.StoragePlugin
		lis net.Listener
	}
	type testCase struct {
		name       string
		fields     fields
		args       args
		before     func(tc *testCase)
		wantErr    bool
		wantErrMsg string
	}
	tests := []testCase{
		{
			name:   "when the service mode is unset",
			fields: fields{},
			args: args{
				ctx: context.Background(),
				sp:  nil,
				lis: nil,
			},
			before: func(tc *testCase) {
				os := mocks.NewMockOSInterface(gomock.NewController(t))
				os.EXPECT().Getenv("X_CSI_NODE_NAME").Times(1).Return("")
				os.EXPECT().Getenv("NODE_NAME").Times(1).Return("worker-1")
				opSys = os

				err := csictx.Setenv(tc.args.ctx, gocsi.EnvVarMode, "")
				if err != nil {
					t.Errorf("failed to set env var for mode. err: %s", err.Error())
				}
			},
			wantErr:    true,
			wantErrMsg: "mode not set",
		},
		{
			name:   "fail to create k8s client",
			fields: fields{},
			args: args{
				ctx: context.Background(),
				sp:  nil,
				lis: nil,
			},
			before: func(tc *testCase) {
				os := mocks.NewMockOSInterface(gomock.NewController(t))
				os.EXPECT().Getenv("X_CSI_NODE_NAME").Times(1).Return("worker-1")
				opSys = os

				// for validateGlobalVars func
				DriverNamespace = "powerstore"
				DriverName = "csi-powerstore.dellemc.com"

				err := csictx.Setenv(tc.args.ctx, gocsi.EnvVarMode, "controller")
				if err != nil {
					t.Errorf("failed to set env var for mode. err: %s", err.Error())
				}
			},
			wantErr:    true,
			wantErrMsg: "unable to load in-cluster configuration",
		},
		{
			name:   "fail to get the node",
			fields: fields{},
			args: args{
				ctx: context.Background(),
				sp:  nil,
				lis: nil,
			},
			before: func(tc *testCase) {
				os := mocks.NewMockOSInterface(gomock.NewController(t))
				os.EXPECT().Getenv("X_CSI_NODE_NAME").Times(1).Return("worker-1")
				opSys = os

				// for validateGlobalVars func
				DriverNamespace = "powerstore"
				DriverName = "csi-powerstore.dellemc.com"

				err := csictx.Setenv(tc.args.ctx, gocsi.EnvVarMode, "controller")
				if err != nil {
					t.Errorf("failed to set env var for mode. err: %s", err.Error())
				}

				clientSet := fake.NewClientset()
				k8s.RestInClusterConfigFunc = func() (*rest.Config, error) {
					return new(rest.Config), nil
				}
				k8s.NewForConfigFunc = func(config *rest.Config) (kubernetes.Interface, error) {
					return clientSet, nil
				}
			},
			wantErr:    true,
			wantErrMsg: "not found",
		},
		{
			name:   "",
			fields: fields{},
			args: args{
				ctx: context.Background(),
				sp:  nil,
				lis: nil,
			},
			before: func(tc *testCase) {
				os := mocks.NewMockOSInterface(gomock.NewController(t))
				os.EXPECT().Getenv("X_CSI_NODE_NAME").Times(1).Return("worker-1")
				opSys = os

				// for validateGlobalVars func
				DriverNamespace = "powerstore"
				DriverName = "csi-powerstore.dellemc.com"

				err := csictx.Setenv(tc.args.ctx, gocsi.EnvVarMode, "controller")
				if err != nil {
					t.Errorf("failed to set env var for mode. err: %s", err.Error())
				}

				clientSet := fake.NewClientset()
				k8s.RestInClusterConfigFunc = func() (*rest.Config, error) {
					return new(rest.Config), nil
				}
				k8s.NewForConfigFunc = func(config *rest.Config) (kubernetes.Interface, error) {
					return clientSet, nil
				}
			},
			wantErr:    true,
			wantErrMsg: "not found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer afterEach()
			tt.before(&tt)

			s := &CsiNfsService{
				vcsi:            tt.fields.vcsi,
				md:              tt.fields.md,
				provisionerName: tt.fields.provisionerName,
				mode:            tt.fields.mode,
				nodeID:          tt.fields.nodeID,
				nodeIPAddress:   tt.fields.nodeIPAddress,
				podCIDR:         tt.fields.podCIDR,
				nodeName:        tt.fields.nodeName,
				k8sclient:       tt.fields.k8sclient,
				executor:        tt.fields.executor,
			}
			err := s.BeforeServe(tt.args.ctx, tt.args.sp, tt.args.lis)
			if (err != nil) != tt.wantErr {
				t.Errorf("CsiNfsService.BeforeServe() error = %v, wantErr %v", err, tt.wantErr)
			}
			// confirm error message matches the expected message.
			if tt.wantErr && (err != nil) {
				if !strings.Contains(err.Error(), tt.wantErrMsg) {
					t.Errorf("CsiNfsService.BeforeServe() error = %v, wantErr %v", err, tt.wantErrMsg)
				}
			}
		})
	}
}

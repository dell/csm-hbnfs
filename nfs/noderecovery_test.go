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
	"testing"

	k8s "github.com/dell/csm-hbnfs/nfs/k8s"
	"github.com/dell/csm-hbnfs/nfs/mocks"
	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "k8s.io/api/core/v1"
	discoveryv1 "k8s.io/api/discovery/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestReassignVolume(t *testing.T) {
	// Create a new fake clientset
	clientset := fake.NewSimpleClientset()

	// Test case: GetPersistentVolume fails
	s := &CsiNfsService{
		k8sclient: &k8s.K8sClient{
			Clientset: clientset,
		},
	}

	slice := &discoveryv1.EndpointSlice{
		ObjectMeta: metav1.ObjectMeta{
			Name: "mySlice",
			Labels: map[string]string{
				"pvName": "pv1",
				"nodeID": "myNode",
			},

			Annotations: map[string]string{
				DriverVolumeID: "vol1",
			},
		},
		Endpoints: []discoveryv1.Endpoint{
			{
				Addresses: []string{"1.2.3.4"},
			},
		},
	}

	if s.reassignVolume(slice) {
		t.Errorf("reassignVolume should return false when GetPersistentVolume fails")
	}

	// Test case: GetService fails
	clientset.CoreV1().PersistentVolumes().Create(context.Background(), &v1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name: "pv1",
		},
		Spec: v1.PersistentVolumeSpec{
			PersistentVolumeSource: v1.PersistentVolumeSource{
				CSI: &v1.CSIPersistentVolumeSource{
					Driver:       "myDriver",
					VolumeHandle: CsiNfsPrefixDash + uuid.New().String(),
				},
			},
		},
	}, metav1.CreateOptions{})

	if s.reassignVolume(slice) {
		t.Errorf("reassignVolume should return false when GetService fails")
	}

	// Test case: callUnexportNfsVolume fails
	clientset.CoreV1().Services("").Create(context.Background(), &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "mySlice",
			Labels: map[string]string{
				"client/myClient": "127.0.0.1",
			},
		},
	}, metav1.CreateOptions{})

	// Set the export counts for the client (will need mux)
	exportCounts["127.0.0.1"] = 2

	service := mocks.NewMockService(gomock.NewController(t))
	service.EXPECT().ControllerUnpublishVolume(gomock.Any(), gomock.Any()).Times(1).Return(nil, status.Errorf(codes.Internal, "unable to unpublish volume"))
	s.vcsi = service

	if s.reassignVolume(slice) {
		t.Errorf("reassignVolume should return false when callUnexportNfsVolume fails")
	}

	// // Test case: ControllerUnpublishVolume fails
	// s = &CsiNfsService{
	// 	k8sclient: &MockK8sClient{
	// 		GetNodeErr: fmt.Errorf("failed to get Node"),
	// 	},
	// 	vcsi: &MockVCSI{
	// 		ControllerUnpublishVolumeErr: fmt.Errorf("failed to unpublish volume"),
	// 	},
	// }
	// if s.reassignVolume(slice) {
	// 	t.Errorf("reassignVolume should return false when ControllerUnpublishVolume fails")
	// }
	// // Test case: GetNode fails
	// s = &CsiNfsService{
	// 	k8sclient: &MockK8sClient{
	// 		GetNodeErr: fmt.Errorf("failed to get Node"),
	// 	},
	// }
	// if s.reassignVolume(slice) {
	// 	t.Errorf("reassignVolume should return false when GetNode fails")
	// }
	// // Test case: ControllerPublishVolume fails
	// s = &CsiNfsService{
	// 	k8sclient: &MockK8sClient{
	// 		GetNodeErr: fmt.Errorf("failed to get Node"),
	// 	},
	// 	vcsi: &MockVCSI{
	// 		ControllerPublishVolumeErr: fmt.Errorf("failed to publish volume"),
	// 	},
	// }
	// if s.reassignVolume(slice) {
	// 	t.Errorf("reassignVolume should return false when ControllerPublishVolume fails")
	// }
	// // Test case: callExportNfsVolume fails
	// s = &CsiNfsService{
	// 	k8sclient: &MockK8sClient{
	// 		GetNodeErr: fmt.Errorf("failed to get Node"),
	// 	},
	// 	vcsi: &MockVCSI{
	// 		ControllerPublishVolumeErr: fmt.Errorf("failed to publish volume"),
	// 	},
	// }
	// if s.reassignVolume(slice) {
	// 	t.Errorf("reassignVolume should return false when callExportNfsVolume fails")
	// }
	// // Test case: UpdateEndpointSlice fails
	// s = &CsiNfsService{
	// 	k8sclient: &MockK8sClient{
	// 		GetNodeErr: fmt.Errorf("failed to get Node"),
	// 	},
	// 	vcsi: &MockVCSI{
	// 		ControllerPublishVolumeErr: fmt.Errorf("failed to publish volume"),
	// 	},
	// }
	// if s.reassignVolume(slice) {
	// 	t.Errorf("reassignVolume should return false when UpdateEndpointSlice fails")
	// }
}

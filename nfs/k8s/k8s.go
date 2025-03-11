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

package k8s

import (
	"context"
	"encoding/json"
	"fmt"

	//	"path/filepath"

	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	discoveryv1 "k8s.io/api/discovery/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	// "k8s.io/client-go/tools/clientcmd"
	// "k8s.io/client-go/util/homedir"
)

type K8sClient struct {
	Clientset *kubernetes.Clientset
}

// Connect connect establishes a connection with the k8s API server.
func Connect() (*K8sClient, error) {
	k8sclient := new(K8sClient)
	var err error
	log.Info("csi-nfs: attempting k8sapi connection using InClusterConfig")
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	k8sclient.Clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Error("csi-nfs: unable to connect to k8sapi: " + err.Error())
		return nil, err
	}
	log.Info("csi-nfs: connected to k8sapi")
	return k8sclient, nil
}

func (kc *K8sClient) CreateService(ctx context.Context, namespace string, service *v1.Service) (*v1.Service, error) {
	return kc.Clientset.CoreV1().Services(namespace).Create(ctx, service, metav1.CreateOptions{})
}

func (kc *K8sClient) GetService(ctx context.Context, namespace, name string) (*v1.Service, error) {
	return kc.Clientset.CoreV1().Services(namespace).Get(ctx, name, metav1.GetOptions{})
}

func (kc *K8sClient) UpdateService(ctx context.Context, namespace string, service *v1.Service) (*v1.Service, error) {
	return kc.Clientset.CoreV1().Services(namespace).Update(ctx, service, metav1.UpdateOptions{})
}

func (kc *K8sClient) DeleteService(ctx context.Context, namespace, name string) error {
	return kc.Clientset.CoreV1().Services(namespace).Delete(ctx, name, metav1.DeleteOptions{})
}

func (kc *K8sClient) GetEndpointSlice(ctx context.Context, namespace, name string) (*discoveryv1.EndpointSlice, error) {
	return kc.Clientset.DiscoveryV1().EndpointSlices(namespace).Get(ctx, name, metav1.GetOptions{})
}

func (kc *K8sClient) CreateEndpointSlice(ctx context.Context, namespace string, endpointSlice *discoveryv1.EndpointSlice) (*discoveryv1.EndpointSlice, error) {
	return kc.Clientset.DiscoveryV1().EndpointSlices(namespace).Create(ctx, endpointSlice, metav1.CreateOptions{})
}

func (kc *K8sClient) UpdateEndpointSlice(ctx context.Context, namespace string, endpointSlice *discoveryv1.EndpointSlice) (*discoveryv1.EndpointSlice, error) {
	return kc.Clientset.DiscoveryV1().EndpointSlices(namespace).Update(ctx, endpointSlice, metav1.UpdateOptions{})
}

func (kc *K8sClient) DeleteEndpointSlice(ctx context.Context, namespace, name string) error {
	return kc.Clientset.DiscoveryV1().EndpointSlices(namespace).Delete(ctx, name, metav1.DeleteOptions{})
}

func (kc *K8sClient) GetPersistentVolume(ctx context.Context, name string) (*v1.PersistentVolume, error) {
	return kc.Clientset.CoreV1().PersistentVolumes().Get(ctx, name, metav1.GetOptions{})
}

func (kc *K8sClient) GetNode(ctx context.Context, nodeName string) (*v1.Node, error) {
	node, err := kc.Clientset.CoreV1().Nodes().Get(ctx, nodeName, metav1.GetOptions{})
	return node, err
}

// GetlEndpointSlices returns the endpointslices matching match labels.
func (k *K8sClient) GetEndpointSlices(ctx context.Context, namespace, labelSelector string) ([]*discoveryv1.EndpointSlice, error) {
	slices := make([]*discoveryv1.EndpointSlice, 0)
	log.Infof("csi-nfs: retrieving all endpointslices")
	sliceList, err := k.Clientset.DiscoveryV1().EndpointSlices(namespace).List(ctx, metav1.ListOptions{LabelSelector: labelSelector})
	if err != nil {
		log.Errorf("csi-nfs: error retrieving endpointslices: %s: %s", labelSelector, err.Error())
	}
	for _, slice := range sliceList.Items {
		slices = append(slices, &slice)
	}
	log.Infof("csi-nfs: retrieved %d endpointslices", len(slices))
	return slices, err
}

// GetAllNodes retrieves all nodes in the Kubernetes cluster
func (k *K8sClient) GetAllNodes(ctx context.Context) ([]*v1.Node, error) {
	nodes := make([]*v1.Node, 0)
	log.Info("csi-nfs: retrieving all nodes")
	nodeList, err := k.Clientset.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Error("csi-nfs: error retrieving nodes: " + err.Error())
		return nil, err
	}
	for _, node := range nodeList.Items {
		nodes = append(nodes, &node)
	}
	log.Infof("csi-nfs: retrieved %d nodes", len(nodeList.Items))
	return nodes, nil
}

func (kc *K8sClient) GetNodeByCSINodeId(ctx context.Context, driverKey string, csiNodeId string) (*v1.Node, error) {
	nodeList, err := kc.Clientset.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to list nodes: %w", err)
	}
	for _, node := range nodeList.Items {
		if annotation, exists := node.Annotations["csi.volume.kubernetes.io/nodeid"]; exists {
			var nodeIdMap map[string]string
			if err := json.Unmarshal([]byte(annotation), &nodeIdMap); err != nil {
				log.Printf("Failed to unmarshal annotation for node %s: %v", node.Name, err)
				continue
			}
			if value, found := nodeIdMap[driverKey]; found && value == csiNodeId {
				return &node, nil
			}
		}
	}
	return nil, fmt.Errorf("failed to find a Node matching csiNodeId %s", csiNodeId)
}

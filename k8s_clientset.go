// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s

import (
	"github.com/golang/glog"
	"github.com/pkg/errors"
	apiextensionsClient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	k8s_kubernetes "k8s.io/client-go/kubernetes"
	k8s_rest "k8s.io/client-go/rest"
	k8s_clientcmd "k8s.io/client-go/tools/clientcmd"
)

func CreateClientset(kubeconfig string) (Interface, error) {
	config, err := CreateConfig(kubeconfig)
	if err != nil {
		return nil, errors.Wrap(err, "create k8s config failed")
	}
	clientset, err := k8s_kubernetes.NewForConfig(config)
	if err != nil {
		return nil, errors.Wrap(err, "build k8s clientset failed")
	}
	return clientset, nil
}

func CreateApiextensionsClient(kubeconfig string) (ApiextensionsInterface, error) {
	config, err := CreateConfig(kubeconfig)
	if err != nil {
		return nil, errors.Wrap(err, "create k8s config failed")
	}
	clientset, err := apiextensionsClient.NewForConfig(config)
	if err != nil {
		return nil, errors.Wrap(err, "build k8s api apiextensions clientset failed")
	}
	return clientset, nil
}

func CreateConfig(kubeconfig string) (*k8s_rest.Config, error) {
	if len(kubeconfig) > 0 {
		glog.V(4).Infof("create kube config from flags")
		return k8s_clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	glog.V(4).Infof("create in cluster kube config")
	return k8s_rest.InClusterConfig()
}

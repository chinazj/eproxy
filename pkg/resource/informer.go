// Copyright (c) 2016-2017 ByteDance, Inc. All rights reserved.
package resource

import (
	"k8s.io/client-go/informers"
	v1 "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"time"
)

const defaultResync = 0 * time.Second

type Resources struct {
	informers       informers.SharedInformerFactory
	ServiceInformer v1.ServiceInformer
	EndpointInfomer v1.EndpointsInformer
}

func NewResources(client kubernetes.Interface) *Resources {
	resources := &Resources{}
	informers := informers.NewSharedInformerFactory(client, defaultResync)
	resources.ServiceInformer = informers.Core().V1().Services()
	resources.EndpointInfomer = informers.Core().V1().Endpoints()
	resources.SetEndpointHandler(&EndpointHandler{})
	resources.SetServiceHandler(&ServiceHandler{})
	return resources
}

func (resources *Resources) SetServiceHandler(handler cache.ResourceEventHandler) {
	resources.ServiceInformer.Informer().AddEventHandler(handler)
}

func (resources *Resources) SetEndpointHandler(handler cache.ResourceEventHandler) {
	resources.EndpointInfomer.Informer().AddEventHandler(handler)
}

func (resources *Resources) StartListenEventFromKubernetes(stopCh <-chan struct{}) {
	resources.ServiceInformer.Informer().Run(stopCh)
	resources.EndpointInfomer.Informer().Run(stopCh)
}

// Copyright (c) 2016-2017 ByteDance, Inc. All rights reserved.
package resource

import "k8s.io/client-go/tools/cache"

type EventType int

const (
	ServiceAdd    EventType = 1
	ServiceDelete EventType = 2
	ServiceUpdate EventType = 3

	EndpointAdd    EventType = 4
	EndpointDelete EventType = 5
	EndpointUpdate EventType = 6
)

type KubernetesEvent struct {
	KType     EventType
	Name      string
	Namespace string
}

type ServiceHandler struct {
	event chan KubernetesEvent
}

func (s *ServiceHandler) OnAdd(obj interface{}) {
	s.event <- KubernetesEvent{}
}

func (s *ServiceHandler) OnUpdate(oldObj, newObj interface{}) {
	s.event <- KubernetesEvent{}
}

func (s *ServiceHandler) OnDelete(obj interface{}) {
	s.event <- KubernetesEvent{}
}

var _ cache.ResourceEventHandler = &ServiceHandler{}

type EndpointHandler struct {
	event chan KubernetesEvent
}

func (s *EndpointHandler) OnAdd(obj interface{}) {
	s.event <- KubernetesEvent{}
}

func (s *EndpointHandler) OnUpdate(oldObj, newObj interface{}) {
	s.event <- KubernetesEvent{}
}

func (s *EndpointHandler) OnDelete(obj interface{}) {
	s.event <- KubernetesEvent{}
}

var _ cache.ResourceEventHandler = &EndpointHandler{}

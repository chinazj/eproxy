// Copyright (c) 2016-2017 ByteDance, Inc. All rights reserved.
package bpf

import "github.com/cilium/ebpf"

type Service struct {
	ipv6   bool
	lb4map ebpf.Map
	lb6map ebpf.Map
	// cache
	service map[ServiceKey]ServiceValue
}

func (s *Service) IsIpv6() bool {
	return s.ipv6
}

func (s *Service) LookUpElemSerivceMap(key ServiceKey) ServiceValue {
	value := Service4Value{}
	s.lb4map.Lookup(key, &value)
	return &value
}

func (s *Service) DeleteElemSerivceMap(Key ServiceKey) error {
	err := s.lb4map.Delete(Key)
	if err == nil {
		delete(s.service, Key)
	}
	return err
}

func (s *Service) UpdateElemSerivceMap(Key ServiceKey, value ServiceValue) error {
	err := s.lb4map.Update(Key, value, ebpf.UpdateAny)
	if err == nil {
		s.service[Key] = value
	}
	return err
}

var _ ServiceMap = &Service{}

type Endpoint struct {
	ipv6            bool
	lb_endpoint_map ebpf.Map
	lb6map          ebpf.Map
	// cache
	service map[EndpointKey]EndpointValue
}

func (e *Endpoint) LookUpElemSerivceMap(key EndpointKey) EndpointValue {
	//TODO implement me
	panic("implement me")
}

func (e *Endpoint) DeleteElemSerivceMap(Key EndpointKey) error {
	//TODO implement me
	panic("implement me")
}

func (e *Endpoint) UpdateElemSerivceMap(Key EndpointKey, value EndpointValue) error {
	//TODO implement me
	panic("implement me")
}

var _ EndpointMap = &Endpoint{}

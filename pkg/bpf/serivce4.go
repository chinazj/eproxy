// Copyright (c) 2016-2017 ByteDance, Inc. All rights reserved.
package bpf

// Service4Key 必须和bpf代码对齐
type Service4Key struct {
	ServiceIP    uint32
	ServicePort  uint16
	Backend_slot uint8
	Proto        uint8
	Node         uint8
	Pad          pad2uint8
}

// Service4Value 必须和bpf代码对齐
type Service4Value struct {
	BackendID uint32
	Count     uint16
	Pad       pad2uint8
}

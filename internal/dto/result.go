package dto

import (
	"sync"
)

type DslResult struct {
	mu           sync.Mutex
	NextNodeName string
	NextCategory string
	Decision     interface{}
	Track        []string
	Detail       []NodeResult
}

type NodeResult struct {
	mu       sync.RWMutex
	NodeName string
	Factor   map[string]Feature
	Hits     []interface{} //hit rules
	Decision interface{}
}

func NewNodeResult(name string) *NodeResult {
	return &NodeResult{NodeName: name, Factor: make(map[string]Feature, 0)}
}

func (ns *NodeResult) AddFactor(factor map[string]Feature) {
	ns.mu.Lock()
	defer ns.mu.Unlock()
	for k, v := range factor {
		ns.Factor[k] = v
	}
}

func (ns *NodeResult) AddHits(rule ...interface{}) {
	ns.mu.Lock()
	defer ns.mu.Unlock()
	ns.Hits = append(ns.Hits, rule...)
}

func (ns *NodeResult) SetDecision(decision interface{}) {
	ns.mu.Lock()
	defer ns.mu.Unlock()
	ns.Decision = decision
}

func NewDslResult() *DslResult {
	return new(DslResult)
}

func (ds *DslResult) AddDetail(ns ...NodeResult) {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	ds.Detail = append(ds.Detail, ns...)
}

func (ds *DslResult) AddTrack(track ...string) {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	ds.Track = append(ds.Track, track...)
}

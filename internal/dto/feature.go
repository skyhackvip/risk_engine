package dto

import (
	"sync"
)

const (
	FeatureTypeInt = iota
	FeatureTypeFloat
	FeatureTypeString
	FeatureTypeBool
	FeatureTypeEnum
)

type Feature struct {
	Name    string
	Type    int
	Value   interface{}
	Default interface{}
}

type GlobalFeatureS struct {
	mu       sync.RWMutex
	Features map[string]Feature
}

func NewGlobalFeatures() *GlobalFeatureS {
	return &GlobalFeatureS{Features: make(map[string]Feature)}
}

func (fs *GlobalFeatureS) Set(f Feature) {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	fs.Features[f.Name] = f
}

func (fs *GlobalFeatureS) Get(names []string) map[string]Feature {
	fs.mu.RLock()
	defer fs.mu.RUnlock()
	var rs = make(map[string]Feature, 0)
	for _, name := range names {
		if v, ok := fs.Features[name]; !ok {
			rs[name] = Feature{}
		} else {
			rs[name] = v
		}
	}
	return rs
}

package main

import (
	"github.com/skyhackvip/risk_engine/dslparser"
	"github.com/skyhackvip/risk_engine/global"
	"github.com/skyhackvip/risk_engine/internal/dto"
	"testing"
)

func init() {
	global.DslResult = dto.NewDslResult()
	global.Features = dto.NewGlobalFeatures()
	features := map[string]interface{}{
		"feature_1": 60,
		"feature_2": 20,
		"feature_3": 30,
		"feature_4": 40,
	}
	for k, v := range features {
		global.Features.Set(dto.Feature{Name: k, Value: v})
	}
}

func TestAbtest(t *testing.T) {
	dsl := dslparser.LoadDslFromFile("../yaml/flow_abtest.yaml")
	dsl.Parse(global.DslResult)
	return
}

package main

import (
	"github.com/skyhackvip/risk_engine/configs"
	"github.com/skyhackvip/risk_engine/dslparser"
	"github.com/skyhackvip/risk_engine/global"
	"github.com/skyhackvip/risk_engine/internal/dto"
	"testing"
)

func init() {
	global.DslResult = dto.NewDslResult()
	global.Features = dto.NewGlobalFeatures()
	features := map[string]interface{}{
		"feature_1": 20,
		"feature_2": 10,
		"feature_3": 10,
		"feature_4": 10,
		"feature_5": 10,
		"feature_6": 10,
	}
	for k, v := range features {
		global.Features.Set(dto.Feature{Name: k, Value: v})
	}
}

func TestFlowSimple(t *testing.T) {
	dsl := dslparser.LoadDslFromFile("../yaml/flow_simple.yaml")
	rs := dsl.Parse(global.DslResult).Decision
	if rs == nil {
		t.Error("nil")
		return
	}
	if rs.(int) == configs.DecisionMap["reject"] {
		t.Log("ok")
	} else {
		t.Log("result error!")
	}
}

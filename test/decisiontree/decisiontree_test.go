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
		"feature_1": 18,
		"feature_2": true,
	}
	for k, v := range features {
		global.Features.Set(dto.Feature{Name: k, Value: v})
	}
}

func TestDecisionTree(t *testing.T) {
	dsl := dslparser.LoadDslFromFile("../yaml/decisiontree.yaml")

	rs, err := dsl.ParseDecisionTree(dsl.DecisionTrees[0])
	if err != nil {
		t.Error(err)
	}
	if rs == "C" {
		t.Log("result is ", rs)
	} else {
		t.Error("result error,result is ", rs)
	}
}

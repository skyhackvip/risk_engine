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
		"model_1": 85,
		"model_2": 180,
	}
	for k, v := range features {
		global.Features.Set(dto.Feature{Name: k, Value: v})
	}
}

func TestDecisionMatrix(t *testing.T) {
	dsl := dslparser.LoadDslFromFile("../yaml/decisionmatrix.yaml")
	rs, err := dsl.ParseDecisionMatrix(dsl.DecisionMatrixs[0])
	if err != nil {
		t.Error(err)
	}
	if rs == "C" {
		t.Log("result is ", rs)
	} else {
		t.Error("result error,expert C, result is ", rs)
	}
}

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
		"feature_1": 60,
		"feature_2": 20,
		"feature_3": 30,
		"feature_4": 40,
	}
	for k, v := range features {
		global.Features.Set(dto.Feature{Name: k, Value: v})
	}
}

func TestConditional(t *testing.T) {

	dsl := dslparser.LoadDslFromFile("../yaml/flow_conditional.yaml")
	result := dsl.Parse(global.DslResult)

	if result.Decision == configs.DecisionMap["reject"] {
		t.Log("Decision result is: pass")
	} else {
		t.Error("Decision result is:", result.Decision)
	}

	t.Log("Decision track is:", result.Track)

}

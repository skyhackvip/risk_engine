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
		"feature_2": 10,
		"feature_3": 10,
	}
	for k, v := range features {
		global.Features.Set(dto.Feature{Name: k, Value: v})
	}
}

func TestRuleset(t *testing.T) {
	dsl := dslparser.LoadDslFromFile("../yaml/ruleset.yaml")
	rs, err := dsl.ParseRuleset(dsl.Rulesets[0], global.DslResult)
	if err != nil {
		t.Error(err)
	}
	if rs == configs.DecisionMap["reject"] {
		t.Log("result is ", rs)
	} else {
		t.Error("result error!")
	}
}

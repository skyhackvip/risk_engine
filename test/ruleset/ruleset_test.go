package main

import (
	"github.com/skyhackvip/risk_engine/configs"
	"github.com/skyhackvip/risk_engine/dslparser"
	"testing"
)

func TestRuleset(t *testing.T) {
	dsl := dslparser.LoadDslFromFile("ruleset.yaml")
	rs, err := dsl.ParseRuleset(dsl.Rulesets[0])
	if err != nil {
		t.Error(err)
	}
	if rs == configs.DecisionMap["reject"] {
		t.Log("result is ", rs)
	} else {
		t.Error("result error!")
	}
}

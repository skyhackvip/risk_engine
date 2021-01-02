package main

import (
	//	"github.com/skyhackvip/risk_engine/configs"
	"github.com/skyhackvip/risk_engine/dslparser"
	"github.com/skyhackvip/risk_engine/internal"
	"testing"
)

func TestDecisionTree(t *testing.T) {
	internal.SetFeature("feature_1", 50)
	internal.SetFeature("feature_2", true)
	dsl := dslparser.LoadDslFromFile("decisiontree.yaml")

	rs, err := dsl.ParseDecisionTree(dsl.DecisionTrees[0])
	if err != nil {
		t.Error(err)
	}
	if rs == "D" {
		t.Log("result is ", rs)
	} else {
		t.Error("result error,expert D, result is ", rs)
	}
}

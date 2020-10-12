package main

import (
	//	"github.com/skyhackvip/risk_engine/configs"
	"github.com/skyhackvip/risk_engine/dslparser"
	"github.com/skyhackvip/risk_engine/internal"
	"testing"
)

func TestDecisionTree(t *testing.T) {
	internal.SetFeature("feature_1", 18)
	internal.SetFeature("feature_2", false)
	dsl := dslparser.LoadDslFromFile("decisiontree.yaml")

	rs := dsl.ParseDecisionTree(dsl.Decisiontrees[0])
	if rs == "D" {
		t.Log("result is ", rs)
	} else {
		t.Error("result error,expert D, result is ", rs)
	}
}

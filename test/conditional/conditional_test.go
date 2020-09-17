package main

import (
	"github.com/skyhackvip/risk_engine/configs"
	"github.com/skyhackvip/risk_engine/dslparser"
	"github.com/skyhackvip/risk_engine/internal"
	"testing"
)

func TestConditional(t *testing.T) {
	internal.SetFeature("feature_1", 10)
	internal.SetFeature("feature_2", 10)
	internal.SetFeature("feature_3", 10)
	internal.SetFeature("feature_4", 10)

	dsl := dslparser.LoadDslFromFile("conditional.yaml")
	result := dsl.Parse()

	if result.Decision == configs.DecisionMap["reject"] {
		t.Log("Decision result is: pass")
	} else {
		t.Error("Decision result is:", result.Decision)
	}

	t.Log("Decision track is:", result.Track)

}

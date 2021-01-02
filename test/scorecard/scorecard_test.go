package main

import (
	//	"github.com/skyhackvip/risk_engine/configs"
	"github.com/skyhackvip/risk_engine/dslparser"
	"github.com/skyhackvip/risk_engine/internal"
	"testing"
)

func TestScoreCard(t *testing.T) {
	internal.SetFeature("amout", 7999)
	internal.SetFeature("sex", "M")
	dsl := dslparser.LoadDslFromFile("scorecard.yaml")

	rs, err := dsl.ParseScoreCard(dsl.ScoreCards[0])
	if err != nil {
		t.Error(err)
	}
	rs = rs.(float64)
	if rs == 7 {
		t.Log("result is ", rs)
	} else {
		t.Error("result error,expert 7, result is ", rs)
	}
}

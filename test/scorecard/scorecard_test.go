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
		"amout": 7999,
		"sex":   "M",
	}
	for k, v := range features {
		global.Features.Set(dto.Feature{Name: k, Value: v})
	}
}

func TestScoreCard(t *testing.T) {
	dsl := dslparser.LoadDslFromFile("../yaml/scorecard.yaml")

	rs, err := dsl.ParseScoreCard(dsl.ScoreCards[0])
	if err != nil {
		t.Error(err)
	}
	rs = rs.(float64)
	if rs == 7.0 {
		t.Log("result is ", rs)
	} else {
		t.Error("result error,expert 7, result is ", rs)
	}
}

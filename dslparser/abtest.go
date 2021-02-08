package dslparser

import (
	"github.com/skyhackvip/risk_engine/global"
	"github.com/skyhackvip/risk_engine/internal/dto"
	"github.com/skyhackvip/risk_engine/internal/errcode"
	"log"
	"math/rand"
	"time"
)

type Abtest struct {
	Name    string   `yaml:"name"`
	Branchs []Branch `yaml:"branchs,flow"`
}

func (abtest *Abtest) parse(result *dto.DslResult) (interface{}, error) {
	rand.Seed(time.Now().UnixNano())
	winNum := rand.Float64() * 100
	var counter float64 = 0
	for _, branch := range abtest.Branchs {
		counter += branch.Percent
		if counter > winNum {
			global.Features.Set(dto.Feature{Name: abtest.Name, Value: branch.BranchName})
			log.Printf("abtest %v : %v, %v\n", abtest.Name, branch.BranchName, winNum)
			return branch.Decision, nil
		}
	}
	return nil, errcode.ParseErrorNoBranchMatch
}

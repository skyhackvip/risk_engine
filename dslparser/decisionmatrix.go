package dslparser

import (
	"fmt"
	"github.com/skyhackvip/risk_engine/global"
	"github.com/skyhackvip/risk_engine/internal/errcode"
	"log"
	"strings"
)

type DecisionMatrix struct {
	Name      string     `yaml:"name"`
	Depends   []string   `yaml:"depends,flow"`
	Rules     []Rule     `yaml:"rules,flow"`
	Decisions []Decision `yaml:"decisions,flow"`
}

func (dm *DecisionMatrix) parse() (interface{}, error) {
	log.Printf("decisionmatrix %s parse ...\n", dm.Name)
	depends := global.Features.Get(dm.Depends)
	log.Println("depend", depends)
	var result = make([]string, 0)
	for _, rule := range dm.Rules {
		rs, err := rule.parse(depends)
		if err != nil {
			return nil, err
		}
		if rs.(bool) { //true will be added
			result = append(result, rule.Decision)
		}
	}
	for _, decision := range dm.Decisions {
		//compare slice []{x,y}
		if compareSlice(decision.Depends, result) {
			return decision.Output, nil
		}
	}
	return nil, errcode.ParseErrorDecisionMatrixOutputEmpty
}

func compareSlice(s1, s2 []string) bool {
	s1Str := strings.Replace(strings.Trim(fmt.Sprint(s1), "[]"), " ", "_", -1)
	s2Str := strings.Replace(strings.Trim(fmt.Sprint(s2), "[]"), " ", "_", -1)
	return s1Str == s2Str
}

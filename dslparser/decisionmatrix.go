package dslparser

import (
	"fmt"
	"github.com/skyhackvip/risk_engine/internal"
	"log"
	"strings"
)

type DecisionMatrix struct {
	Name      string     `yaml:"name"`
	Depends   []string   `yaml:"depends,flow"`
	Rules     []Rule     `yaml:"rules,flow"`
	Decisions []Decision `yaml:"decisions,flow"`
}

func (dm *DecisionMatrix) parse() string {
	log.Printf("decisionmatrix %s parse ...\n", dm.Name)
	depends := internal.GetFeatures(dm.Depends)
	log.Println("depend", depends)
	var result = make([]string, 0)
	for _, rule := range dm.Rules {
		if rule.parse() { //true will be added
			result = append(result, rule.Decision)
		}
	}
	for _, decision := range dm.Decisions {
		//compare slice []{x,y}
		if compareSlice(decision.Depends, result) {
			return decision.Output
		}
	}
	return ""
}

func compareSlice(s1, s2 []string) bool {
	s1Str := strings.Replace(strings.Trim(fmt.Sprint(s1), "[]"), " ", "_", -1)
	s2Str := strings.Replace(strings.Trim(fmt.Sprint(s2), "[]"), " ", "_", -1)
	return s1Str == s2Str
}

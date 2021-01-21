package dslparser

import (
	"github.com/skyhackvip/risk_engine/global"
	"github.com/skyhackvip/risk_engine/internal/errcode"
	"github.com/skyhackvip/risk_engine/internal/operator"
	"log"
)

type DecisionTree struct {
	Name      string     `yaml:"name"`
	Depends   []string   `yaml:"depends,flow"`
	Rules     []Rule     `yaml:"rules,flow"`
	Decisions []Decision `yaml:"decisions,flow"`
}

func (dt *DecisionTree) parse() (interface{}, error) {
	log.Printf("decisiontree %s parse ...\n", dt.Name)
	var result = make(map[string]bool, 0)
	depends := global.Features.Get(dt.Depends)
	for _, rule := range dt.Rules {
		rs, err := rule.parse(depends)
		if err != nil {
			return nil, err
		}
		result[rule.Decision] = rs.(bool)
	}
	for _, decision := range dt.Decisions {
		if parseDecision(result, decision) {
			return decision.Output, nil
		}
	}
	return nil, errcode.ParseErrorDecisionTreeOutputEmpty
}

func parseDecision(result map[string]bool, decision Decision) bool {
	var rs = make([]bool, 0)
	for _, depend := range decision.Depends {
		if data, ok := result[depend]; ok {
			rs = append(rs, data)
		}
	}
	final, _ := operator.Boolean(rs, decision.Logic)
	return final
}

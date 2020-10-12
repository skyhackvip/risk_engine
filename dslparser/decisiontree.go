package dslparser

import (
	//	"github.com/skyhackvip/risk_engine/internal"
	"github.com/skyhackvip/risk_engine/operator"
	"log"
)

type Decisiontree struct {
	Name      string     `yaml:"name"`
	Depends   []string   `yaml:"depends,flow"`
	Rules     []Rule     `yaml:"rules,flow"`
	Decisions []Decision `yaml:"decisions,flow"`
}

func (dt *Decisiontree) parse() string {
	log.Printf("decisiontree %s parse ...\n", dt.Name)
	var result = make(map[string]bool, 0)
	for _, rule := range dt.Rules {
		result[rule.Decision] = rule.parse()
	}
	for _, decision := range dt.Decisions {
		if parseDecision(result, decision) {
			return decision.Output
		}
	}
	return ""
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

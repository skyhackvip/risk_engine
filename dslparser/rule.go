package dslparser

import (
	"github.com/skyhackvip/risk_engine/internal"
	"github.com/skyhackvip/risk_engine/operator"
	"log"
)

type Rule struct {
	Conditions []Condition `yaml:"conditions,flow"`
	RuleName   string      `yaml:"rule_name"`
	RuleGroup  string      `yaml:"rule_group"`
	Logic      string      `yaml:"logic"`
	Decision   string      `yaml:"decision"`
	Depends    []string    `yaml:"depends"`
}

//parse rule
func (rule *Rule) parse() bool {
	var conditionRs = make([]bool, 0)
	depends := internal.GetFeatures(rule.Depends) //need to check
	for _, condition := range rule.Conditions {
		if data, ok := depends[condition.Feature]; ok {
			rs, _ := operator.Compare(condition.Operator, data, condition.Value)
			log.Printf("rule %s parse : %v %v %v ,result is: %v\n", rule.RuleName, data, condition.Operator, condition.Value, rs)
			conditionRs = append(conditionRs, rs)
		}
	}
	logicRs, _ := operator.Boolean(conditionRs, rule.Logic)
	log.Printf("rule %s decision is: %v\n", rule.RuleName, logicRs)
	return logicRs
}

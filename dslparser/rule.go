package dslparser

import (
	//"github.com/skyhackvip/risk_engine/global"
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
func (rule *Rule) parse() (interface{}, error) {
	var conditionRs = make([]bool, 0)
	depends := internal.GetFeatures(rule.Depends) //need to check
	for _, condition := range rule.Conditions {
		if data, ok := depends[condition.Feature]; ok {
			rs, err := operator.Compare(condition.Operator, data, condition.Value)
			log.Printf("rule %s parse : %v %v %v ,result is: %v\n", rule.RuleName, data, condition.Operator, condition.Value, rs)
			if err != nil {
				return nil, err
			}
			conditionRs = append(conditionRs, rs)
		}
	}
	logicRs, err := operator.Boolean(conditionRs, rule.Logic)
	log.Printf("rule %s decision is: %v\n", rule.RuleName, logicRs)
	if err != nil {
		return nil, err
	}
	return logicRs, err
}

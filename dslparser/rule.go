package dslparser

import (
	"github.com/skyhackvip/risk_engine/configs"
	"github.com/skyhackvip/risk_engine/internal"
	"github.com/skyhackvip/risk_engine/operator"
	"log"
)

type Rule struct {
	Conditions []Condition `yaml:"conditions,flow"`
	RuleName   string      `yaml:"rule_name"`
	Logic      string      `yaml:"logic"`
	Decision   string      `yaml:"decision"`
	Depends    []string    `yaml:"depends"`
}

//parse rule
func (rule *Rule) parse() int {
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
	if logicRs {
		return configs.DecisionMap[rule.Decision]
	} else {
		return configs.NilDecision
	}
}

//get depend feature value
func getDepends(depends []string) map[string]interface{} {
	var rs = make(map[string]interface{})
	rsMap := make(map[string]int)
	rsMap["feature1"] = 5
	rsMap["feature2"] = 10
	rsMap["feature3"] = 20

	for _, feature := range depends {
		rs[feature] = rsMap[feature]
	}
	return rs
}

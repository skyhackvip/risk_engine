package dslparser

import (
	"github.com/skyhackvip/risk_engine/internal/dto"
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
func (rule *Rule) parse(depends map[string]dto.Feature) (interface{}, error) {
	var conditionRs = make([]bool, 0)

	for _, condition := range rule.Conditions {
		if data, ok := depends[condition.Feature]; ok {
			if data.Name == "" { //TODO
				log.Println("data error")
				continue
			}
			rs, err := operator.Compare(condition.Operator, data.Value, condition.Value)
			//log.Printf("rule %s parse : %v %v %v ,result is: %v\n", rule.RuleName, data.Value, condition.Operator, condition.Value, rs)
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

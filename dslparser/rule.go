package dslparser

import (
	"errors"
	"fmt"
	"github.com/skyhackvip/risk_engine/internal/dto"
	"github.com/skyhackvip/risk_engine/internal/operator"
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
	if len(rule.Conditions) == 0 {
		return nil, errors.New(fmt.Sprintf("rule (%s) condition is empty", rule.RuleName))
	}
	for _, condition := range rule.Conditions {
		if data, ok := depends[condition.Feature]; ok {
			if data.Name == "" {
				log.Println("data error : data name is empty")
				continue
			}
			rs, err := operator.Compare(condition.Operator, data.Value, condition.Value)
			if err != nil {
				return nil, err
			}
			conditionRs = append(conditionRs, rs)
		} else {
			//lack of feature
			log.Printf("lack of feature:%s\n", condition.Feature)
			continue
		}
	}
	if len(conditionRs) == 0 {
		return nil, errors.New(fmt.Sprintf("rule (%s) condition is empty", rule.RuleName))
	}
	logicRs, err := operator.Boolean(conditionRs, rule.Logic)
	log.Printf("rule %s decision is: %v\n", rule.RuleName, logicRs)
	if err != nil {
		return nil, err
	}
	return logicRs, nil
}

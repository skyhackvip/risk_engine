package dslparser

import (
	"github.com/skyhackvip/risk_engine/internal"
	"github.com/skyhackvip/risk_engine/internal/errcode"
	"github.com/skyhackvip/risk_engine/operator"
)

//conditinal gateway node
type Conditional struct {
	ConditionalName string   `yaml:"conditional_name"`
	Depends         []string `yaml:"depends"`
	Branchs         []Branch `yaml:"branchs,flow"`
}

//branch in conditional
type Branch struct {
	BranchName string      `yaml:"branch_name"`
	Conditions []Condition `yaml:"conditions"`
	Logic      string      `yaml:"logic"`
	Decision   string      `yaml:"decision"`
}

//conditional gateway parse
func (conditional *Conditional) parse() (interface{}, error) {
	depends := internal.GetFeatures(conditional.Depends) //need to check
	for _, branch := range conditional.Branchs {         //loop all the branch
		var conditionRs = make([]bool, 0)
		for _, condition := range branch.Conditions {
			if data, ok := depends[condition.Feature]; ok {
				rs, err := operator.Compare(condition.Operator, data, condition.Value)
				if err != nil {
					return nil, err
				}
				conditionRs = append(conditionRs, rs)
			} else { //get feature fail
				continue //can modify according scene
			}
		}
		logicRs, _ := operator.Boolean(conditionRs, branch.Logic)
		if logicRs { //if true, choose the branch and break
			return branch.Decision, nil
		} else {
			continue
		}
	}
	return nil, errcode.ParseErrorNoBranchMatch //can't find any branch
}

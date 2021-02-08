package dslparser

import (
	"github.com/skyhackvip/risk_engine/global"
	"github.com/skyhackvip/risk_engine/internal/dto"
	"github.com/skyhackvip/risk_engine/internal/errcode"
	"github.com/skyhackvip/risk_engine/internal/operator"
)

//conditinal gateway node
type Conditional struct {
	ConditionalName string   `yaml:"conditional_name"`
	Depends         []string `yaml:"depends"`
	Branchs         []Branch `yaml:"branchs,flow"`
}

//conditional gateway parse
func (conditional *Conditional) parse(result *dto.DslResult) (interface{}, error) {
	nodeResult := dto.NewNodeResult(conditional.ConditionalName)
	depends := global.Features.Get(conditional.Depends)
	nodeResult.AddFactor(depends)
	for _, branch := range conditional.Branchs { //loop all the branch
		var conditionRs = make([]bool, 0)
		for _, condition := range branch.Conditions {
			if data, ok := depends[condition.Feature]; ok {
				//data
				//TODO
				rs, err := operator.Compare(condition.Operator, data.Value, condition.Value)
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
			nodeResult.SetDecision(branch.Decision)
			result.AddDetail(*nodeResult)
			return branch.Decision, nil
		} else {
			continue
		}
	}
	return nil, errcode.ParseErrorNoBranchMatch //can't find any branch
}

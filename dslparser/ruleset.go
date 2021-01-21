package dslparser

import (
	"github.com/skyhackvip/risk_engine/configs"
	"github.com/skyhackvip/risk_engine/global"
	"github.com/skyhackvip/risk_engine/internal/dto"
	"github.com/skyhackvip/risk_engine/internal/errcode"
	"log"
	"sort"
)

type Ruleset struct {
	RulesetName     string   `yaml:"ruleset_name"`
	RulesetCategory string   `yaml:"ruleset_category"`
	RuleExec        string   `yaml:"rule_exec"`
	Rules           []Rule   `yaml:"rules,flow"`
	Depends         []string `yaml:"depends,flow"`
}

func (ruleset *Ruleset) parse(result *dto.DslResult) (interface{}, error) {
	log.Printf("----- ruleset %s parse -------\n", ruleset.RulesetName)
	nodeResult := dto.NewNodeResult(ruleset.RulesetName)
	var ruleResult = make([]int, 0)
	depends := global.Features.Get(ruleset.Depends)
	nodeResult.AddFactor(depends)
	for _, rule := range ruleset.Rules {
		rs, err := rule.parse(depends)
		if err != nil {
			return nil, err
		}
		ruleDecision := configs.NilDecision
		if rs.(bool) { //HIT
			nodeResult.Hits = append(nodeResult.Hits, rule.RuleName)
			ruleDecision = configs.DecisionMap[rule.Decision]
		}
		ruleResult = append(ruleResult, ruleDecision)
	}
	if len(ruleResult) == 0 {
		log.Printf("ruleset %s parse no result\n", ruleset.RulesetName)
		return nil, errcode.ParseErrorRulesetOutputEmpty
	}
	//get max value result, reject is 100, record is 1, pass or no result is 0
	sort.Sort(sort.Reverse(sort.IntSlice(ruleResult)))
	log.Printf("ruleset %s result is :%v\n", ruleset.RulesetName, ruleResult[0])
	nodeResult.Decision = ruleResult[0]
	result.AddDetail(*nodeResult)
	log.Printf("----- ruleset %s end -------\n", ruleset.RulesetName)
	return ruleResult[0], nil
}

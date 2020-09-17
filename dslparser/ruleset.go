package dslparser

import (
	"log"
	"sort"
)

type Ruleset struct {
	RulesetName     string `yaml:"ruleset_name"`
	RulesetCategory string `yaml:"ruleset_category"`
	RuleExec        string `yaml:"rule_exec"`
	Rules           []Rule `yaml:"rules,flow"`
}

func (ruleset *Ruleset) parse() int {
	log.Printf("ruleset %s parse :\n", ruleset.RulesetName)
	var ruleResult = make([]int, 0)
	for _, rule := range ruleset.Rules {
		ruleResult = append(ruleResult, rule.parse())
	}
	if len(ruleResult) == 0 {
		log.Printf("ruleset %s parse no result\n", ruleset.RulesetName)
	}
	//get max value result, reject is 100, record is 1, pass or no result is 0
	sort.Sort(sort.Reverse(sort.IntSlice(ruleResult)))
	log.Printf("ruleset %s result is :%v\n", ruleset.RulesetName, ruleResult[0])
	return ruleResult[0]
}

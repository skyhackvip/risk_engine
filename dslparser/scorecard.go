package dslparser

import (
	"github.com/skyhackvip/risk_engine/configs"
	"github.com/skyhackvip/risk_engine/global"
	"github.com/skyhackvip/risk_engine/internal/operator"
	"log"
	"strconv"
	"strings"
)

type ScoreCard struct {
	Name     string   `yaml:"name"`
	Depends  []string `yaml:"depends,flow"`
	Rules    []Rule   `yaml:"rules,flow"`
	Decision Decision `yaml:"decision"`
}

func (sc *ScoreCard) parse() (interface{}, error) {
	log.Printf("scorecard %s parse ...\n", sc.Name)
	var result = make(map[string]string, 0)
	depends := global.Features.Get(sc.Depends)
	for _, rule := range sc.Rules {
		if _, exists := result[rule.RuleGroup]; !exists {
			rs, err := rule.parse(depends)
			if err != nil {
				return nil, err
			}
			if rs.(bool) { //hit
				result[rule.RuleGroup] = rule.Decision
			}
		}
	}
	var scores = make([]string, 0)
	for _, v := range result {
		scores = append(scores, v)
	}
	return parseScoreCard(scores, sc.Decision.Logic, sc.Decision.Output)
}

func parseScoreCard(scores []string, logic string, output string) (interface{}, error) {
	var score float64
	switch logic {
	case configs.Sum:
		scoreStr, _ := operator.Math(strings.Join(scores, "+"))
		score = scoreStr.(float64)
	}
	expr := strings.Replace(output, configs.ScoreReplace, strconv.FormatFloat(score, 'f', -1, 64), -1)
	result, err := operator.Math(expr)
	if err != nil {
		return nil, err
	}
	return result, nil
}

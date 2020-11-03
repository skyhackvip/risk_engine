package dslparser

import (
	"github.com/skyhackvip/risk_engine/configs"
	"github.com/skyhackvip/risk_engine/operator"
	"log"
	"strconv"
	"strings"
)

type ScoreCard struct {
	Name     string   `yaml:"name"`
	Depends  string   `yaml:"depends"`
	Rules    []Rule   `yaml:"rules,flow"`
	Decision Decision `yaml:"decision"`
}

func (sc *ScoreCard) parse() float64 {
	log.Printf("scorecard %s parse ...\n", sc.Name)
	var result = make(map[string]string, 0)
	for _, rule := range sc.Rules {
		if _, exists := result[rule.RuleGroup]; !exists {
			if rule.parse() { //hit
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

func parseScoreCard(scores []string, logic string, output string) float64 {
	var score float64
	switch logic {
	case configs.Sum:
		scoreStr, _ := operator.Math(strings.Join(scores, "+"))
		score = scoreStr.(float64)
	}
	expr := strings.Replace(output, configs.ScoreReplace, strconv.FormatFloat(score, 'f', -1, 64), -1)
	result, _ := operator.Math(expr)
	return result.(float64)
}

package dslparser

import (
	"github.com/skyhackvip/risk_engine/configs"
	"github.com/skyhackvip/risk_engine/internal/dto"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Dsl struct {
	Workflow        []Node           `yaml:"workflow,flow"`
	Rulesets        []Ruleset        `yaml:"rulesets,flow"`
	Conditionals    []Conditional    `yaml:"conditionals,flow"`
	Abtests         []Abtest         `yaml:"abtests,flow"`
	DecisionTrees   []DecisionTree   `yaml:"decisiontrees,flow"`
	DecisionMatrixs []DecisionMatrix `yaml:"decisionmatrixs,flow"`
	ScoreCards      []ScoreCard      `yaml:"scorecards,flow"`
}

//load dsl from file
func LoadDslFromFile(file string) *Dsl {
	dsl := new(Dsl)
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, dsl)

	if err != nil {
		panic(err)
	}
	return dsl
}

//parse dsl run node followed workflow
func (dsl *Dsl) Parse(result *dto.DslResult) *dto.DslResult {
	log.Println("dsl parse start...")
	if len(dsl.Workflow) == 0 {
		panic("dsl workflow is empty!")
	}
	//from start node
	firstNode := dsl.FindStartNode()
	dsl.gotoNextNode(firstNode.NodeName, firstNode.Category, result)

	//loop parse node and go to next node
	for result.NextNodeName != "" && !isBreakDecision(result.Decision) {
		dsl.gotoNextNode(result.NextNodeName, result.NextCategory, result)
	}
	log.Println("dsl parse end!")
	return result
}

//if decision is break decision (reject),break and exit workflow
func isBreakDecision(decision interface{}) bool {
	if decision == nil {
		return false
	}
	return decision.(int) == configs.DecisionMap[configs.BreakDecision]
}

//parse node and find next
func (dsl *Dsl) gotoNextNode(nodeName string, category string, result *dto.DslResult) {
	//find current node from workflow
	node := dsl.FindNode(nodeName)
	if node == nil {
		return
	}
	result.Track = append(result.Track, nodeName)
	//default
	result.NextNodeName = node.NextNodeName
	result.NextCategory = node.NextCategory
	result.Decision = nil

	//parse different category node
	switch category {
	case configs.START:
		return
	case configs.RULESET:
		ruleset := dsl.FindRuleset(node.NodeName)
		result.Decision, _ = ruleset.parse(result)
	case configs.CONDITIONAL:
		conditional := dsl.FindConditional(node.NodeName)
		rs, _ := conditional.parse(result)
		if rs == nil { //not match any branch, error
			result.NextNodeName = ""
			log.Println(node.NodeName, "not match any branch")
		} else {
			result.NextNodeName = rs.(string)
			result.NextCategory = dsl.FindNode(rs.(string)).Category
		}
	case configs.ABTEST:
		abtest := dsl.FindAbtest(node.NodeName)
		rs, _ := abtest.parse(result)
		if rs == nil { //not match any branch, error
			result.NextNodeName = ""
			log.Println(node.NodeName, "not match any branch")
		} else {
			result.NextNodeName = rs.(string)
			result.NextCategory = dsl.FindNode(rs.(string)).Category
		}
	case configs.DECISIONTREE:
		decisionTree := dsl.FindDecisionTree(node.NodeName)
		rs, _ := decisionTree.parse()
		result.Decision = rs
	case configs.DECISIONMATRIX:
		decisionMatrix := dsl.FindDecisionMatrix(node.NodeName)
		rs, _ := decisionMatrix.parse()
		result.Decision = rs
	case configs.SCORECARD:
		scorecard := dsl.FindScoreCard(node.NodeName)
		rs, _ := scorecard.parse()
		result.Decision = rs
	case configs.END:
		result.NextNodeName = ""
		result.NextCategory = ""
	}
}

//dsl.Rulesets []Ruleset
func (dsl *Dsl) FindRuleset(name string) *Ruleset {
	for _, ruleset := range dsl.Rulesets {
		if ruleset.RulesetName == name {
			return &ruleset
		}
	}
	return nil
}

//dsl.Conditionals []Condtional
func (dsl *Dsl) FindConditional(name string) *Conditional {
	for _, conditional := range dsl.Conditionals {
		if conditional.ConditionalName == name {
			return &conditional
		}
	}
	return nil
}

//dsl.Abtests []Abtest
func (dsl *Dsl) FindAbtest(name string) *Abtest {
	for _, abtest := range dsl.Abtests {
		if abtest.Name == name {
			return &abtest
		}
	}
	return nil
}

//dsl.DecisionTrees []DecisionTree
func (dsl *Dsl) FindDecisionTree(name string) *DecisionTree {
	for _, decisionTree := range dsl.DecisionTrees {
		if decisionTree.Name == name {
			return &decisionTree
		}
	}
	return nil
}

//dsl.DecisionMatrixs []DecisionMatrix
func (dsl *Dsl) FindDecisionMatrix(name string) *DecisionMatrix {
	for _, decisionMatrix := range dsl.DecisionMatrixs {
		if decisionMatrix.Name == name {
			return &decisionMatrix
		}
	}
	return nil
}

//dsl.ScoreCards []ScoreCard
func (dsl *Dsl) FindScoreCard(name string) *ScoreCard {
	for _, scoreCard := range dsl.ScoreCards {
		if scoreCard.Name == name {
			return &scoreCard
		}
	}
	return nil
}

//dsl.Workflow []Node
//find nodename=name
func (dsl *Dsl) FindNode(name string) *Node {
	for _, node := range dsl.Workflow {
		if node.NodeName == name {
			return &node
		}
	}
	return nil
}

//dsl.Workflow []Node
//category = start
func (dsl *Dsl) FindStartNode() *Node {
	for _, node := range dsl.Workflow {
		if node.Category == configs.START {
			return &node
		}
	}
	return nil
}

//parse ruleset
func (dsl *Dsl) ParseRuleset(ruleset Ruleset, result *dto.DslResult) (interface{}, error) {
	log.Println("test parse ruleset")
	return ruleset.parse(result)
}

//parse conditional
func (dsl *Dsl) ParseConditional(conditional Conditional, result *dto.DslResult) (interface{}, error) {
	return conditional.parse(result)
}

//parse abtest
func (dsl *Dsl) ParseAbtest(abtest Abtest, result *dto.DslResult) (interface{}, error) {
	return abtest.parse(result)
}

//parse decisiontree
func (dsl *Dsl) ParseDecisionTree(decisionTree DecisionTree) (interface{}, error) {
	return decisionTree.parse()
}

//parse decisionmatrix
func (dsl *Dsl) ParseDecisionMatrix(decisionMatrix DecisionMatrix) (interface{}, error) {
	return decisionMatrix.parse()
}

//parse scorecard
func (dsl *Dsl) ParseScoreCard(sc ScoreCard) (interface{}, error) {
	return sc.parse()
}

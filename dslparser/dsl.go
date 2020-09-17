package dslparser

import (
	"github.com/skyhackvip/risk_engine/configs"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Dsl struct {
	Workflow     []Node        `yaml:"workflow,flow"`
	Rulesets     []Ruleset     `yaml:"rulesets,flow"`
	Conditionals []Conditional `yaml:"conditionals,flow"`
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
func (dsl *Dsl) Parse() *DslResult {
	log.Println("dsl parse start...")
	if len(dsl.Workflow) == 0 {
		panic("dsl workflow is empty!")
	}
	var result = new(DslResult)
	//from start node
	firstNode := dsl.FindStartNode()
	dsl.gotoNextNode(firstNode.NodeName, firstNode.Category, result)

	//loop parse node and go to next node
	for !isBreakDecision(result.Decision) && result.NextNodeName != "" {
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
func (dsl *Dsl) gotoNextNode(nodeName string, category string, result *DslResult) {
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
		result.Decision = ruleset.parse()
	case configs.CONDITIONAL:
		conditional := dsl.FindConditional(node.NodeName)
		rs := conditional.parse()
		if rs == "" { //not match any branch, error
			result.NextNodeName = ""
			log.Println(node.NodeName, "not match any branch")
		} else {
			result.NextNodeName = rs
			result.NextCategory = dsl.FindNode(rs).Category
		}
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
func (dsl *Dsl) ParseRuleset(ruleset Ruleset) interface{} {
	return ruleset.parse()
}

//parse conditional
func (dsl *Dsl) ParseConditional(conditional Conditional) interface{} {
	return conditional.parse()
}

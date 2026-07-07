// Copyright (c) 2023
//
// @author 贺鹏Kavin
// 微信公众号:技术岁月
// https://github.com/skyhackvip/risk_engine
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
package core

import (
	"github.com/skyhackvip/risk_engine/internal/log"
)

type Dsl struct {
	Key          string                 `yaml:"key"`
	Version      string                 `yaml:"version"`
	Metadata     map[string]interface{} `yaml:"metadata"`
	DecisionFlow []FlowNode             `yaml:"decision_flow,flow"`
	Features     []Feature              `yaml:"features,flow"`
	Abtests      []AbtestNode           `yaml:"abtests,flow"`
	Conditionals []ConditionalNode      `yaml:"conditionals,flow"`
	Rulesets     []RulesetNode          `yaml:"rulesets,flow"`
	Matrixs      []MatrixNode           `yaml:"matrixs,flow"`
	Trees        []TreeNode             `yaml:"trees,flow"`
	Scorecards   []ScorecardNode        `yaml:"scorecards,flow"`
}

func (dsl *Dsl) CheckValid() bool {
	if dsl.Key == "" {
		return false
	}
	if len(dsl.DecisionFlow) == 0 {
		return false
	}
	return true
}

//dsl to decisionflow
func (dsl *Dsl) ConvertToDecisionFlow() (*DecisionFlow, error) {
	flow := NewDecisionFlow()
	flow.Key = dsl.Key
	flow.Version = dsl.Version
	flow.Metadata = dsl.Metadata

	//map
	featureMap := make(map[string]IFeature)
	for _, feature := range dsl.Features {
		featureMap[feature.Name] = NewFeature(feature.Name, GetFeatureType(feature.Kind)) //IFeature
	}
	flow.FeatureMap = featureMap
	rulesetMap := make(map[string]INode)
	for _, ruleset := range dsl.Rulesets {
		rulesetMap[ruleset.GetName()] = ruleset
	}
	abtestMap := make(map[string]INode)
	for _, abtest := range dsl.Abtests {
		abtestMap[abtest.GetName()] = abtest
	}
	conditionalMap := make(map[string]INode)
	for _, conditional := range dsl.Conditionals {
		conditionalMap[conditional.GetName()] = conditional
	}
	matrixMap := make(map[string]INode)
	for _, martix := range dsl.Matrixs {
		matrixMap[martix.GetName()] = martix
	}
	treeMap := make(map[string]INode)
	for _, tree := range dsl.Trees {
		treeMap[tree.GetName()] = tree
	}
	scorecardMap := make(map[string]INode)
	for _, scorecard := range dsl.Scorecards {
		scorecardMap[scorecard.GetName()] = scorecard
	}

	//flow
	for _, flowNode := range dsl.DecisionFlow {
		newNode := flowNode //need set new variable
		switch GetNodeType(newNode.NodeKind) {
		case TypeRuleset:
			newNode.SetElem(rulesetMap[newNode.NodeName])
			flow.AddNode(&newNode)
		case TypeAbtest:
			newNode.SetElem(abtestMap[newNode.NodeName])
			flow.AddNode(&newNode)
		case TypeConditional:
			newNode.SetElem(conditionalMap[newNode.NodeName])
			flow.AddNode(&newNode)
		case TypeStart:
			newNode.SetElem(NewStartNode(newNode.NodeName))
			flow.SetStartNode(&newNode)
			flow.AddNode(&newNode)
		case TypeEnd:
			newNode.SetElem(NewEndNode(newNode.NodeName))
			flow.AddNode(&newNode)
		case TypeMatrix:
			newNode.SetElem(matrixMap[newNode.NodeName])
			flow.AddNode(&newNode)
		case TypeTree:
			newNode.SetElem(treeMap[newNode.NodeName])
			flow.AddNode(&newNode)
		case TypeScorecard:
			newNode.SetElem(scorecardMap[newNode.NodeName])
			flow.AddNode(&newNode)
		default:
			log.Warnf("dsl %s - %s convert warning: unknown node type %s", dsl.Key, dsl.Version, newNode.NodeKind)
		}
	}
	return flow, nil
}

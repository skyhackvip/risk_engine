package operator

import (
	"errors"
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/skyhackvip/risk_engine/configs"
)

//compare expression:left [><=] right
func Compare(operator string, left interface{}, right interface{}) (bool, error) {
	var params = make(map[string]interface{})
	params["left"] = left
	params["right"] = right
	var expr *govaluate.EvaluableExpression
	if _, ok := configs.OperatorMap[operator]; !ok {
		return false, errors.New("not support operator")
	}
	expr, _ = govaluate.NewEvaluableExpression(fmt.Sprintf("left %s right", configs.OperatorMap[operator]))
	eval, err := expr.Evaluate(params)
	if err != nil {
		return false, err
	}
	if result, ok := eval.(bool); ok {
		return result, nil
	}
	return false, errors.New("convert error")
}

package operator

import (
	"errors"
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/skyhackvip/risk_engine/configs"
)

//boolean logic expresion : true [&& ||] false
func Boolean(result []bool, logic string) (bool, error) {
	resultLen := len(result)
	if resultLen == 0 {
		return false, errors.New("at least 1 input")
	}
	if resultLen == 1 {
		return result[0], nil
	}
	var exprStr string
	for i := 0; i < resultLen; i++ {
		exprStr += fmt.Sprintf(" %t", result[i])
		if i != (resultLen - 1) {
			exprStr += fmt.Sprintf(" %s", configs.LogicMap[logic])
		}
	}
	expr, _ := govaluate.NewEvaluableExpression(exprStr)
	eval, err := expr.Evaluate(nil)
	if err != nil {
		return false, err
	}
	if result, ok := eval.(bool); ok {
		return result, nil
	}
	return false, errors.New("convert error")
}

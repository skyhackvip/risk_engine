package operator

import (
	"github.com/Knetic/govaluate"
)

func Math(expression string) (interface{}, error) {
	expr, _ := govaluate.NewEvaluableExpression(expression)
	result, err := expr.Evaluate(nil)
	return result, err
}

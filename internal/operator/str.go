package operator

import "errors"

//support eq,contain,like
func StrLogic(operator string, left string, right string) (bool, error) {
	switch operator {
	case "EQ":
		return left == right, nil
	}
	return false, errors.New("not support operator")
}

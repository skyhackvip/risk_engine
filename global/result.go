package global

type DslResult struct {
	NextNodeName string
	NextCategory string
	Decision     interface{}
	Track        []string
	Detail       []interface{}
}

var rs DslResult

func SetResult(rs DslResult) {
	rs = rs
}

func GetResult() DslResult {
	return rs
}

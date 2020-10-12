package global

type DslResult struct {
	NextNodeName string
	NextCategory string
	Decision     interface{}
	Track        []string
	Detail       []interface{}
}

package dslparser

type BaseNode interface {
	parse() (interface{}, error)
}

package dslparser

type Base interface {
	parse() (interface{}, error)
}

package dslparser

type Condition struct {
	Feature  string      `yaml:"feature"`
	Operator string      `yaml:"operator"`
	Value    interface{} `yaml:"value"`
	Result   string      `yaml:"result"`
}

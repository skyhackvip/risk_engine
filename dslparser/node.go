package dslparser

type Node struct {
	NodeName     string `yaml:"node_name"`
	Category     string `yaml:"category"`
	NextNodeName string `yaml:"next_node_name"`
	NextCategory string `yaml:"next_category"`
}

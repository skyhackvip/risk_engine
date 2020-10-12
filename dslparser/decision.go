package dslparser

type Decision struct {
	Depends []string `yaml:"depends,flow"`
	Logic   string   `yaml:"logic"`
	Output  string   `yaml:"output"`
}

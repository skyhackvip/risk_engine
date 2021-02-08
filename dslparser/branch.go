package dslparser

type Branch struct {
	BranchName string      `yaml:"branch_name"`
	Conditions []Condition `yaml:"conditions"` //used by conditional
	Logic      string      `yaml:"logic"`      //used by conditional
	Percent    float64     `yaml:"percent"`    //used by abtest
	Decision   string      `yaml:"decision"`
}

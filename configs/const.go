package configs

var LogicMap = map[string]string{
	"OR":  "||",
	"AND": "&&",
}

var OperatorMap = map[string]string{
	"GT":  ">",
	"LT":  "<",
	"GE":  ">=",
	"LE":  "<=",
	"EQ":  "==",
	"NEQ": "!=",
	//todo add string like
}

var DecisionMap = map[string]int{
	"reject": 100, //first priority
	"pass":   0,
	"record": 1,
}

const (
	ScoreReplace = "((score))"
)

const (
	Sum = "SUM"
	Min = "MIN"
	Max = "MAX"
	Avg = "AVG"
)

//decision
const (
	NilDecision   = 0        //not hit rules strategy
	BreakDecision = "reject" //if hit,break at once
)

//all support node
const (
	START        = "start"
	END          = "end"
	RULESET      = "ruleset"
	ABTEST       = "abtest"
	CONDITIONAL  = "conditional"
	DECISIONTREE = "decisiontree"
)

package errcode

var (
	ParseErrorNoBranchMatch             = NewError(1000005, "flow error, no branch match")
	ParseErrorRulesetOutputEmpty        = NewError(1000011, "ruleset output is empty")
	ParseErrorDecisionTreeOutputEmpty   = NewError(1000012, "decision tree output is empty")
	ParseErrorDecisionMatrixOutputEmpty = NewError(1000013, "decision matrix output is empty")
)

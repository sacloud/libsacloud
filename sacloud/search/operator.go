package search

// ComparisonOperator フィルター比較演算子
type ComparisonOperator string

const (
	// OpEqual =
	OpEqual = ComparisonOperator("")
	// OpGreaterThan >
	OpGreaterThan = ComparisonOperator(">")
	// OpGreaterEqual >=
	OpGreaterEqual = ComparisonOperator(">=")
	// OpLessThan <
	OpLessThan = ComparisonOperator("<")
	// OpLessEqual <=
	OpLessEqual = ComparisonOperator("<=")
)

// LogicalOperator フィルター論理演算子
type LogicalOperator int

const (
	// OpAnd AND
	OpAnd LogicalOperator = iota
	// OpOr OR
	OpOr
)

package search

import (
	"encoding/json"
	"fmt"
	"strings"
)

// EqualExpression Equalで比較する際の条件
type EqualExpression struct {
	Op         LogicalOperator
	Conditions []interface{}
}

// AndEqual 部分一致(Partial Match)かつAND条件を示すEqualFilterを作成
func AndEqual(conditions ...string) *EqualExpression {
	var values []interface{}
	for _, p := range conditions {
		values = append(values, p)
	}

	return &EqualExpression{
		Op:         OpAnd,
		Conditions: values,
	}
}

// OrEqual 完全一致(Partial Match)かつOR条件を示すEqualFilterを作成
func OrEqual(conditions ...interface{}) *EqualExpression {
	return &EqualExpression{
		Op:         OpOr,
		Conditions: conditions,
	}
}

// MarshalJSON .
func (eq *EqualExpression) MarshalJSON() ([]byte, error) {
	var strConds []string
	for _, cond := range eq.Conditions {
		if cond != nil {
			strConds = append(strConds, convertToValidFilterCondition(cond))
		}
	}

	var value interface{}
	switch eq.Op {
	case OpOr:
		value = strConds
	case OpAnd:
		value = strings.Join(strConds, "%20")
	default:
		return nil, fmt.Errorf("invalid search.LogicalOperator: %v", eq.Op)
	}

	return json.Marshal(value)
}

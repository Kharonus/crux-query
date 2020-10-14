package model

import (
	"fmt"
	"strings"

	"olympos.io/encoding/edn"
)

type WhereClause struct {
	EntityKey   string
	PropertyKey edn.Keyword
	Value       *QueryValue
}

func (clause WhereClause) String() string {
	if clause.EntityKey == "" {
		clause.EntityKey = "_"
	}

	return fmt.Sprintf(`[%s %s %s]`, clause.EntityKey, clause.PropertyKey, clause.Value)
}

type KeyList []string

func (keys KeyList) String() string {
	return fmt.Sprintf("[%s]", strings.Join(keys, " "))
}

type QueryValueType int

const (
	Variable QueryValueType = iota
	Constant
)

type QueryValue struct {
	Type QueryValueType
	Key  string
}

func (value *QueryValue) String() string {
	if value == nil {
		return "_"
	}

	switch value.Type {
	case Variable:
		return fmt.Sprintf(`?%s`, value.Key)
	case Constant:
		return fmt.Sprintf(`"%s"`, value.Key)
	default:
		return value.Key
	}
}

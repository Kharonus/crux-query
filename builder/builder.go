package builder

import (
	"fmt"
	"olympos.io/encoding/edn"
	"strings"

	"github.com/Kharonus/crux-query/model"
)

type CruxQueryBuilder interface {
	Build() string
	Where(key, property, value string, valueType model.QueryValueType) CruxQueryBuilder
	WhereT(clause *model.WhereClause) CruxQueryBuilder
	Find(key string) CruxQueryBuilder
	IncludeFullResults(include bool) CruxQueryBuilder
}

type builder struct {
	clauses     []*model.WhereClause
	keysToFind  model.KeyList
	fullResults bool
}

func NewCruxQueryBuilder() CruxQueryBuilder {
	return &builder{}
}

func (builder *builder) Build() string {
	var built string

	if len(builder.keysToFind) > 0 {
		built += fmt.Sprintf(":find %s ", builder.keysToFind.String())
	}

	if len(builder.clauses) > 0 {
		built += fmt.Sprintf(":where %s ", builder.createWhereClausesString())
	}

	if builder.fullResults {
		built += ":full-results? true"
	}

	return strings.TrimSpace(built)
}

func (builder *builder) createWhereClausesString() string {
	var clauses []string

	for _, clause := range builder.clauses {
		clauses = append(clauses, clause.String())
	}

	return fmt.Sprintf("[%s]", strings.Join(clauses, " "))
}

func (builder *builder) Find(key string) CruxQueryBuilder {
	builder.keysToFind = append(builder.keysToFind, key)

	return builder
}

func (builder *builder) Where(key, property, value string, valueType model.QueryValueType) CruxQueryBuilder {
	builder.clauses = append(builder.clauses, &model.WhereClause{
		EntityKey:   key,
		PropertyKey: edn.Keyword(property),
		Value:       &model.QueryValue{Type: valueType, Key: value},
	})

	return builder
}

func (builder *builder) WhereT(clause *model.WhereClause) CruxQueryBuilder {
	builder.clauses = append(builder.clauses, clause)

	return builder
}

func (builder *builder) IncludeFullResults(include bool) CruxQueryBuilder {
	builder.fullResults = include

	return builder
}

package test

import (
	"log"
	"testing"

	"github.com/Kharonus/crux-query/builder"
	. "github.com/Kharonus/crux-query/model"
	"github.com/stretchr/testify/assert"
)

func Test_Build_BuildsExpectedResult(t *testing.T) {
	for _, test := range builderTestCases {
		log.Print(test.description)
		query := test.builder.Build()

		assert.Equal(t, test.expectedQuery, query)
	}
}

type builderTestSuite struct {
	description   string
	builder       builder.CruxQueryBuilder
	expectedQuery string
}

var builderTestCases = []*builderTestSuite{
	{
		description:   "must return empty result on new builder",
		builder:       builder.NewCruxQueryBuilder(),
		expectedQuery: "",
	},
	{
		description: "must return correct query for a builder with single find and simple where clause with constant",
		builder: builder.NewCruxQueryBuilder().
			Find("entity").
			WhereT(&WhereClause{
				EntityKey:   "entity",
				PropertyKey: "crux.db/id",
				Value:       &QueryValue{Type: Constant, Key: "anything"},
			}),
		expectedQuery: `:find [entity] :where [[entity :crux.db/id "anything"]]`,
	},
	{
		description: "must return correct query for a builder with single find and constructed where clause with constant",
		builder: builder.NewCruxQueryBuilder().
			Find("entity").
			Where("entity", "crux.db/id", "anything", Constant),
		expectedQuery: `:find [entity] :where [[entity :crux.db/id "anything"]]`,
	},
	{
		description: "must return correct query for a builder with single find, constructed where clause with constant and full results",
		builder: builder.NewCruxQueryBuilder().
			Find("entity").
			Where("entity", "crux.db/id", "anything", Constant).
			IncludeFullResults(true),
		expectedQuery: `:find [entity] :where [[entity :crux.db/id "anything"]] :full-results? true`,
	},
}

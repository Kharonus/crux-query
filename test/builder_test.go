package test

import (
	"testing"

	"github.com/Kharonus/crux-query/builder"
	"github.com/stretchr/testify/assert"
)

func Test_Build_ReturnsExpectedResult(t *testing.T) {
	// Arrange
	b := builder.NewCruxQueryBuilder()

	// Act
	query := b.Build()

	// Assert
	assert.Empty(t, query)
}

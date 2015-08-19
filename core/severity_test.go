package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSeverityComparison(t *testing.T) {
	assert := assert.New(t)

	comparator := MakeStringSeverityComparator([]string{"LOWEST", "TRACE", "DEBUG", "HIGHEST"})

	assert.True(comparator.eq("TRACE", "TRACE"))
	assert.False(comparator.lt("TRACE", "TRACE"))
	assert.False(comparator.gt("TRACE", "TRACE"))

	assert.True(comparator.lt("TRACE", "DEBUG"))
	assert.False(comparator.gt("TRACE", "DEBUG"))
	assert.False(comparator.eq("TRACE", "DEBUG"))

	assert.True(comparator.gt("DEBUG", "TRACE"))
	assert.False(comparator.eq("DEBUG", "TRACE"))
	assert.False(comparator.lt("DEBUG", "TRACE"))

	assert.True(comparator.lt("LOWEST", "TRACE"))
	assert.True(comparator.lt("LOWEST", "DEBUG"))
	assert.True(comparator.lt("LOWEST", "HIGHEST"))

	assert.True(comparator.gt("HIGHEST", "TRACE"))
	assert.True(comparator.gt("HIGHEST", "DEBUG"))
	assert.True(comparator.gt("HIGHEST", "LOWEST"))
}

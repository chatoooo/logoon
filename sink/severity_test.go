package sink

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
	assert.True(comparator.ge("TRACE", "TRACE"))

	assert.True(comparator.lt("TRACE", "DEBUG"))
	assert.False(comparator.gt("TRACE", "DEBUG"))
	assert.False(comparator.eq("TRACE", "DEBUG"))
	assert.False(comparator.ge("TRACE", "DEBUG"))

	assert.True(comparator.gt("DEBUG", "TRACE"))
	assert.False(comparator.eq("DEBUG", "TRACE"))
	assert.False(comparator.lt("DEBUG", "TRACE"))
	assert.True(comparator.ge("DEBUG", "TRACE"))

	assert.True(comparator.lt("LOWEST", "TRACE"))
	assert.True(comparator.lt("LOWEST", "DEBUG"))
	assert.True(comparator.lt("LOWEST", "HIGHEST"))
	assert.False(comparator.gt("LOWEST", "HIGHEST"))

	assert.True(comparator.gt("HIGHEST", "TRACE"))
	assert.True(comparator.gt("HIGHEST", "DEBUG"))
	assert.True(comparator.gt("HIGHEST", "LOWEST"))
	assert.True(comparator.gt("HIGHEST", "LOWEST"))

	//unknowns are always false
	assert.False(comparator.lt("DEBUG", "UNKNOWN"))
	assert.False(comparator.gt("DEBUG", "UNKNOWN"))
	assert.False(comparator.eq("DEBUG", "UNKNOWN"))
	assert.False(comparator.ge("DEBUG", "UNKNOWN"))

}

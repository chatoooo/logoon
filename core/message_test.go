package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicMessageMethods(t *testing.T) {
	assert := assert.New(t)

	message := CreateBasicLogMessage("SEVERITY", "TEST")
	assert.EqualValues("TEST", message.Message())
	assert.EqualValues("SEVERITY", message.Severity())
	assert.NotNil(message.Time())
}

func TestTaggedMessageMethods(t *testing.T) {
	assert := assert.New(t)

	message := CreateTaggedLogMessage("SEVERITY", "TEST", []string{"ACCESS"})
	assert.EqualValues("TEST", message.Message())
	assert.EqualValues([]string{"ACCESS"}, message.Tags())
	assert.EqualValues("SEVERITY", message.Severity())
	assert.NotNil(message.Time())
}

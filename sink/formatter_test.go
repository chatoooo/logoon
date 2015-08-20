package sink

import (
	"fmt"
	"github.com/chatoooo/logoon/core"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFormatterSpecialDateFormat(t *testing.T) {
	assert := assert.New(t)

	formatter := new(StandardSinkFormatter)
	formatter.SetFormat("[%time(2006-02-01 15:03:04.000)%] [%severity%] - %message%")
	var msg core.LogMessage = core.CreateBasicLogMessage("DEBUG", "test")

	output := formatter.GetFormattedMessage(msg)
	expected := fmt.Sprintf("[%s] [%s] - %s", msg.Time().Format("2006-02-01 15:03:04.000"), "DEBUG", "test")

	assert.EqualValues(expected, output)
}

func TestFormatterDefaultDateFormat(t *testing.T) {
	assert := assert.New(t)

	formatter := new(StandardSinkFormatter)
	formatter.SetFormat("[%time%] [%severity%] - %message%")
	var msg core.LogMessage = core.CreateBasicLogMessage("DEBUG", "test")

	output := formatter.GetFormattedMessage(msg)
	expected := fmt.Sprintf("[%s] [%s] - %s", msg.Time().Format(time.RFC3339), "DEBUG", "test")

	assert.EqualValues(expected, output)
}

func TestFormatterNewLines(t *testing.T) {
	assert := assert.New(t)

	formatter := new(StandardSinkFormatter)
	formatter.SetFormat("[%time%] [%severity%] - %message%")
	var msg core.LogMessage = core.CreateBasicLogMessage("DEBUG", "test\ntest")

	output := formatter.GetFormattedMessage(msg)
	expected := fmt.Sprintf("[%s] [%s] - %s", msg.Time().Format(time.RFC3339), "DEBUG", "test\n\ttest")

	assert.EqualValues(expected, output)
}

func TestFormatterTagged(t *testing.T) {
	assert := assert.New(t)

	formatter := new(StandardSinkFormatter)
	formatter.SetFormat("[%time%] [%severity%] - %message% (%tags%)")
	var msg core.LogMessage = core.CreateTaggedLogMessage("DEBUG", "test", []string{"A", "B"})

	output := formatter.GetFormattedMessage(msg)
	expected := fmt.Sprintf("[%s] [%s] - %s (%s)", msg.Time().Format(time.RFC3339), "DEBUG", "test", "A,B")

	assert.EqualValues(expected, output)
}

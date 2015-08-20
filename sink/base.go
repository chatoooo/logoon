package sink

import (
	"github.com/chatoooo/logoon/core"
)

type StandardSink struct {
	Filter    core.SinkFilter
	Formatter core.SinkFormatter
	Output    core.SinkOutput
	Name      string
}

func BuildStandardSink(name string, filter core.SinkFilter, format core.SinkFormatter, output core.SinkOutput) core.Sink {
	return &StandardSink{
		Filter:    filter,
		Formatter: format,
		Output:    output,
		Name:      name,
	}
}

func (this StandardSink) GetName() string {
	return this.Name
}

func (this StandardSink) Log(msg core.LogMessage) {
	if this.Filter.ShouldOutput(msg) {
		var logMessage string = this.Formatter.GetFormattedMessage(msg)
		this.Output.Write(logMessage)
	}
}

package sink

import (
	"fmt"
	"github.com/chatoooo/logoon/core"
)

type ConsoleSinkOutput struct{}

func (this *ConsoleSinkOutput) Write(value string) {
	fmt.Println(value)
}

func ConsoleSinkFactory(config *core.SinkConfig, severity []string) (core.Sink, error) {
	var formatter core.SinkFormatter = new(StandardSinkFormatter)
	formatter.SetFormat(config.Format)

	var filter core.SinkFilter = &StandardSinkFilter{MakeStringSeverityComparator(severity), nil}
	filter.SetFilter(&config.ParsedFilter)
	output := new(ConsoleSinkOutput)
	return BuildStandardSink(config.Name, filter, formatter, output), nil
}

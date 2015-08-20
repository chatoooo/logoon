package sink

import (
	"github.com/chatoooo/logoon/core"
	"os"
)

type FileSinkOutput struct {
	file *os.File
}

func (this *FileSinkOutput) Write(value string) {
	this.file.WriteString(value)
	this.file.WriteString("\n")
}

func FileSinkFactory(config *core.SinkConfig, severity []string) (core.Sink, error) {
	var formatter core.SinkFormatter = new(StandardSinkFormatter)
	var err error
	formatter.SetFormat(config.Format)

	var filter core.SinkFilter = &StandardSinkFilter{MakeStringSeverityComparator(severity), nil}
	filter.SetFilter(&config.ParsedFilter)
	output := new(FileSinkOutput)
	output.file, err = os.OpenFile(config.Output, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	return BuildStandardSink(config.Name, filter, formatter, output), nil
}

package sink

import (
	"github.com/chatoooo/logoon/core"
	"os"
	"errors"
	//"fmt"
)

type FileSinkOutput struct {
	file *os.File
}

func (this *FileSinkOutput) Write(value string) {
	this.file.WriteString(value)
	this.file.WriteString("\n")
}

func FileSinkFactory(config *core.SinkConfig, severity []string) (core.Sink, error) {
	//fmt.Printf("Creating sink: %#v\n\n", config.ParsedFilter)
	var formatter core.SinkFormatter
	var err error
	var filter core.SinkFilter
	var options map[string]interface{}
	var filename string
	var value interface{}
	var ok bool

	if options , ok = config.Options.(map[string]interface{}); !ok {
		return nil, errors.New("Wrong 'options' format, expected json object")
	}
	if value, ok = options["filename"]; !ok {
		return nil, errors.New("Missing 'filename' option")
	}
	if filename, ok = value.(string); !ok {
		return nil, errors.New("Option 'filename' must be string")
	}

	formatter = new(StandardSinkFormatter)
	formatter.SetFormat(config.Format)

	filter = &StandardSinkFilter{MakeStringSeverityComparator(severity), core.Filter{}}
	filter.SetFilter(config.ParsedFilter)

	output := new(FileSinkOutput)
	output.file, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)

	if err != nil {
		return nil, err
	}

	return BuildStandardSink(config.Name, filter, formatter, output), nil
}

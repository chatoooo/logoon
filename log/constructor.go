package log

import (
	"encoding/json"
	"github.com/chatoooo/logoon/core"
	"github.com/chatoooo/logoon/sink"
	"github.com/chatoooo/logoon/source"
	"io"
	"os"
	"fmt"
)

type SinkFactory func(*core.SinkConfig, []string) (core.Sink, error)

type Logger source.LogSource
type FormatLogger source.LogSourceFormatted

var sinkFactories map[string]SinkFactory
var globalDispatcher core.Dispatcher

func init() {
	sinkFactories = make(map[string]SinkFactory, 2)
	RegisterSinkFactory("file", sink.FileSinkFactory)
	RegisterSinkFactory("console", sink.ConsoleSinkFactory)
}

func CreateLoggingFromFile(configfile string) (core.Dispatcher, error) {
	file, err := os.Open(configfile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return CreateLoggingFromReader(file)
}

func CreateLoggingFromReader(configFileReader io.Reader) (core.Dispatcher, error) {
	var config core.Config
	err := json.NewDecoder(configFileReader).Decode(&config)
	if err != nil {
		return nil, err
	}
	return CreateLogging(&config)
}

func CreateLogging(config *core.Config) (core.Dispatcher, error) {
	var dispatcher core.Dispatcher = core.CreateDispatcher()
	var sink core.Sink
	var err error
	for _, sinkConfig := range config.Sinks{
		if sinkConfig.Disabled {
			continue
		}

		if factory, ok := sinkFactories[sinkConfig.Type]; ok {
			sinkConfig.ParsedFilter = core.ParseFilterConfig(&sinkConfig.Filters)
			sink, err = factory(sinkConfig, config.Severities)
			if err != nil {
				if config.Debug {
					fmt.Printf("ERROR: Unable to create sink type '%s' named '%s': %s\n", sinkConfig.Type, sinkConfig.Name, err.Error())
				}
				continue
			}
			if config.Debug {
				fmt.Printf("INFO: Created sink type '%s' named '%s'\n", sinkConfig.Type, sinkConfig.Name)
			}
			dispatcher.AddSink(sink)
		} else {
			if config.Debug {
				fmt.Printf("WARNING: Unable to find factory for sink type %s\n", sinkConfig.Type)
			}
		}
	}

	return dispatcher, nil
}

func RegisterSinkFactory(name string, factory SinkFactory) {
	sinkFactories[name] = factory
}

func SetGlobalDispatcher(dispatcher core.Dispatcher) {
	globalDispatcher = dispatcher
}

func CreateSimpleLogger(dispatcher core.Dispatcher) Logger {
	if dispatcher == nil {
		dispatcher = globalDispatcher
	}
	return source.MakeSimpleLogSource(dispatcher)
}

func CreateTaggedLogger(dispatcher core.Dispatcher, tags []string) Logger {
	if dispatcher == nil {
		dispatcher = globalDispatcher
	}
	return source.MakeTaggedLogSource(dispatcher, tags)
}

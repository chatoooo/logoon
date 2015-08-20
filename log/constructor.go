package log

import (
	"encoding/json"
	"github.com/chatoooo/logoon/core"
	"github.com/chatoooo/logoon/sink"
	"github.com/chatoooo/logoon/source"
	"io"
	"os"
)

type SinkFactory func(*core.SinkConfig, []string) core.Sink

type Logger source.LogSource
type FormatLogger source.LogSourceFormatted

var sinkFactories map[string]SinkFactory
var globalDispatcher core.Dispatcher

func init() {
	RegisterSinkFactory("file", sink.FileSinkFactory)
	RegisterSinkFactory("console", sink.ConsoleSinkFactory)
}

func CreateDispatcherFromFile(configfile string) (core.Dispatcher, error) {
	file, err := os.Open(configfile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return CreateDispatcherFromReader(file)
}

func CreateDispatcherFromReader(configFileReader io.Reader) (core.Dispatcher, error) {
	var config core.Config
	err := json.NewDecoder(configFileReader).Decode(&config)
	if err != nil {
		return nil, err
	}
	return CreateDispatcher(config)
}

func CreateDispatcher(config core.Config) (core.Dispatcher, error) {
	return nil, nil
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

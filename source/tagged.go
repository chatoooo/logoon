package source

import (
	"github.com/chatoooo/logoon/core"
)

type TaggedLogSource struct {
	dispatcher core.Dispatcher
	tags       []string
}

func MakeTaggedLogSource(dispatcher core.Dispatcher, tags []string) LogSource {
	return &TaggedLogSource{dispatcher: dispatcher, tags: tags}
}

func (this TaggedLogSource) Trace(message string) {
	this.dispatcher.Log(core.CreateTaggedLogMessage("TRACE", message, this.tags))
}

func (this TaggedLogSource) Debug(message string) {
	this.dispatcher.Log(core.CreateTaggedLogMessage("DEBUG", message, this.tags))
}

func (this TaggedLogSource) Warning(message string) {
	this.dispatcher.Log(core.CreateTaggedLogMessage("WARNING", message, this.tags))
}

func (this TaggedLogSource) Error(message string) {
	this.dispatcher.Log(core.CreateTaggedLogMessage("ERROR", message, this.tags))
}

func (this TaggedLogSource) Fatal(message string) {
	this.dispatcher.Log(core.CreateTaggedLogMessage("FATAL", message, this.tags))
}

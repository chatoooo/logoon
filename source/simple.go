package source

import (
	"github.com/chatoooo/logoon/core"
)

type SimpleLogSource struct {
	dispatcher core.Dispatcher
}

func MakeSimpleLogSource(dispatcher core.Dispatcher) LogSource {
	return &SimpleLogSource{dispatcher: dispatcher}
}

func (this SimpleLogSource) Trace(message string) {
	this.dispatcher.Log(core.CreateBasicLogMessage("TRACE", message))
}

func (this SimpleLogSource) Debug(message string) {
	this.dispatcher.Log(core.CreateBasicLogMessage("DEBUG", message))
}

func (this SimpleLogSource) Warning(message string) {
	this.dispatcher.Log(core.CreateBasicLogMessage("WARNING", message))
}

func (this SimpleLogSource) Error(message string) {
	this.dispatcher.Log(core.CreateBasicLogMessage("ERROR", message))
}

func (this SimpleLogSource) Fatal(message string) {
	this.dispatcher.Log(core.CreateBasicLogMessage("FATAL", message))
}

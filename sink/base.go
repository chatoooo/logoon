package sink

import (
	"github.com/chatoooo/logoon/core"
)

type SinkBase struct {
	Name   string
	Filter *core.Filter
}

func (this SinkBase) GetName() string {
	return this.Name
}

func (this SinkBase) ShouldOutput(message core.LogMessage) bool {
	var tagsOK, severityOK bool
	if this.Filter == nil {
		return true
	}

	return tagsOK && severityOK
}

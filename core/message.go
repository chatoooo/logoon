package core

import (
	"time"
)

type LogMessage interface {
	Severity() string
	Tags() []string
	Time() time.Time
	Message() string
	//Params() map[string]string
}

type BasicLogMessage struct {
	severity string
	message  string
	created  time.Time
}

func CreateBasicLogMessage(severity string, message string) LogMessage {
	return &BasicLogMessage{severity: severity, message: message, created: time.Now()}
}

func (this *BasicLogMessage) Severity() string {
	return this.severity
}

func (this *BasicLogMessage) Message() string {
	return this.message
}

func (this *BasicLogMessage) Time() time.Time {
	return this.created
}

func (this *BasicLogMessage) Tags() []string {
	return nil
}

type TaggedLogMessage struct {
	severity string
	message  string
	created  time.Time
	tags     []string
}

func CreateTaggedLogMessage(severity string, message string, tags []string) LogMessage {
	return &TaggedLogMessage{severity: severity, message: message, created: time.Now(), tags: tags}
}

func (this *TaggedLogMessage) Severity() string {
	return this.severity
}

func (this *TaggedLogMessage) Message() string {
	return this.message
}

func (this *TaggedLogMessage) Time() time.Time {
	return this.created
}

func (this *TaggedLogMessage) Tags() []string {
	return this.tags
}

package core

type Dispatcher interface {
	AddSink(Sink)
	Log(LogMessage)
}

type LogDispatcher struct {
	sinks        map[string]Sink
	sinkChannels map[string]chan LogMessage
	comparator   SeverityComparator
}

func CreateDispatcher(severityOrder []string) Dispatcher {
	return &LogDispatcher{
		sinks:        make(map[string]Sink),
		sinkChannels: make(map[string]chan LogMessage),
		comparator:   MakeStringSeverityComparator(severityOrder),
	}
}

func (this LogDispatcher) Log(message LogMessage) {
	for _, v := range this.sinkChannels {
		v <- message
	}
}

func (this LogDispatcher) AddSink(s Sink) {
	var name string = s.GetName()
	var channel = make(chan LogMessage)
	this.sinks[name] = s
	this.sinkChannels[name] = channel
	go s.Run(channel, this.comparator)
}

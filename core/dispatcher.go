package core

type Dispatcher interface {
	AddSink(Sink)
	Log(LogMessage)
}

type LogDispatcher struct {
	sinks        map[string]Sink
	sinkChannels map[string]chan LogMessage
}

func CreateDispatcher() Dispatcher {
	return &LogDispatcher{
		sinks:        make(map[string]Sink),
		sinkChannels: make(map[string]chan LogMessage),
	}
}

func (this LogDispatcher) Log(message LogMessage) {
	for _, v := range this.sinkChannels {
		//fmt.Printf("Sending Message: %#v\nToChannel: %#v\n\n", message, v)
		v <- message
	}
}

func sinkRunner(s Sink, ch <-chan LogMessage) {
	var msg LogMessage
	for{
		msg = <- ch
		//fmt.Printf("Got Message: %#v\n\n%#v\n\n", msg, s)
		s.Log(msg)
	}
}

func (this LogDispatcher) AddSink(s Sink) {
	var name string = s.GetName()
	var channel = make(chan LogMessage, 10)
	this.sinks[name] = s
	this.sinkChannels[name] = channel
	go sinkRunner(s, channel)
}

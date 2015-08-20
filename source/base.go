package source

type LogSource interface {
	Trace(string)
	Debug(string)
	Info(string)
	Warning(string)
	Error(string)
	Fatal(string)
}

type LogSourceFormatted interface {
	Tracef(string, interface{})
	Debugf(string, interface{})
	Infof(string, interface{})
	Warningf(string, interface{})
	Errorf(string, interface{})
	Fatalf(string, interface{})
}

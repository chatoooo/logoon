package source

type LogSource interface {
	Trace(string)
	Debug(string)
	Warning(string)
	Error(string)
	Fatal(string)
}

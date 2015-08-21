package core

const (
	FILTER_CMP_NONE = iota
	FILTER_CMP_GT
	FILTER_CMP_LT
	FILTER_CMP_GE
	FILTER_CMP_EQ
)

type FilterSeverity struct {
	Level string
	CmpOp int
}

type Filter struct {
	Severity    FilterSeverity
	Tags        []string
	TagsExclude bool
}

type Sink interface {
	GetName() string
	Log(LogMessage)
}

type SinkFilter interface {
	SetFilter(Filter)
	ShouldOutput(LogMessage) bool
}

type SinkFormatter interface {
	SetFormat(string)
	GetFormattedMessage(LogMessage) string
}

type SinkOutput interface {
	Write(string)
}

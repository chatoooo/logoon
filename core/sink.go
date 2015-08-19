package core

const (
	FILTER_CMP_EQ = iota
	FILTER_CMP_GT
	FILTER_CMP_LT
)

type Filter struct {
	Severity struct {
		Severity string
		CmpOp    int
	}
	Tags []string
}

type Sink interface {
	SetFilter(*Filter)
	GetName() string
	Run(chan LogMessage, SeverityComparator)
}

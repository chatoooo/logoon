package core

type Severity interface{}

type SeverityComparator interface {
	eq(Severity, Severity) bool
	lt(Severity, Severity) bool
	gt(Severity, Severity) bool
}

type StringSeverityComparator struct {
	SeverityOrder map[string]int
}

func MakeStringSeverityComparator(severityOrder []string) SeverityComparator {
	var comparator = StringSeverityComparator{
		SeverityOrder: make(map[string]int),
	}
	for k, v := range severityOrder {
		comparator.SeverityOrder[v] = k
	}
	return comparator
}

func (this StringSeverityComparator) eq(a, b Severity) bool {
	var valA, valB string
	valA = a.(string)
	valB = b.(string)
	return this.SeverityOrder[valA] == this.SeverityOrder[valB]
}

func (this StringSeverityComparator) lt(a, b Severity) bool {
	var valA, valB string
	valA = a.(string)
	valB = b.(string)
	return this.SeverityOrder[valA] < this.SeverityOrder[valB]
}

func (this StringSeverityComparator) gt(a, b Severity) bool {
	var valA, valB string
	valA = a.(string)
	valB = b.(string)
	return this.SeverityOrder[valA] > this.SeverityOrder[valB]
}

package sink

type Severity interface{}

type SeverityComparator interface {
	ge(Severity, Severity) bool
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
	var okA, okB bool
	var sA, sB int

	valA = a.(string)
	valB = b.(string)
	sA, okA = this.SeverityOrder[valA]
	sB, okB = this.SeverityOrder[valB]

	return okA && okB && sA == sB
}

func (this StringSeverityComparator) lt(a, b Severity) bool {
	var valA, valB string
	var okA, okB bool
	var sA, sB int

	valA = a.(string)
	valB = b.(string)
	sA, okA = this.SeverityOrder[valA]
	sB, okB = this.SeverityOrder[valB]

	return okA && okB && sA < sB
}

func (this StringSeverityComparator) gt(a, b Severity) bool {
	var valA, valB string
	var okA, okB bool
	var sA, sB int

	valA = a.(string)
	valB = b.(string)
	sA, okA = this.SeverityOrder[valA]
	sB, okB = this.SeverityOrder[valB]

	return okA && okB && sA > sB
}

func (this StringSeverityComparator) ge(a, b Severity) bool {
	var valA, valB string
	var okA, okB bool
	var sA, sB int

	valA = a.(string)
	valB = b.(string)
	sA, okA = this.SeverityOrder[valA]
	sB, okB = this.SeverityOrder[valB]

	return okA && okB && sA >= sB
}

package sink

import (
	"github.com/chatoooo/logoon/core"
	//"fmt"
)

type StandardSinkFilter struct {
	severityComparator SeverityComparator
	filter             core.Filter
}

func (this *StandardSinkFilter) SetFilter(filter core.Filter) {
	this.filter = filter
	//fmt.Printf("Setting filter: %#v,\n%#v\n\n", filter, this.filter)
}

func (this *StandardSinkFilter) shouldSeverity(msg core.LogMessage) bool {
	var result bool
	switch this.filter.Severity.CmpOp {
	case core.FILTER_CMP_NONE:
		result = true
	case core.FILTER_CMP_GE:
		result = this.severityComparator.ge(msg.Severity(), this.filter.Severity.Level)
	case core.FILTER_CMP_EQ:
		result = this.severityComparator.eq(msg.Severity(), this.filter.Severity.Level)
	case core.FILTER_CMP_LT:
		result = this.severityComparator.lt(msg.Severity(), this.filter.Severity.Level)
	case core.FILTER_CMP_GT:
		result = this.severityComparator.gt(msg.Severity(), this.filter.Severity.Level)
	}
	return result
}

func (this *StandardSinkFilter) shouldTags(msg core.LogMessage) bool {
	if this.filter.Tags != nil && len(this.filter.Tags) > 0 {
		msgTags := msg.Tags()
		for _, filterTag := range this.filter.Tags {
			for _, messageTag := range msgTags {
				if filterTag == messageTag {
					return true
				}
			}
		}
		return false
	}
	return true
}

func (this *StandardSinkFilter) ShouldOutput(msg core.LogMessage) bool {
	var severity, tags bool

	severity = this.shouldSeverity(msg)
	if !severity {
		return severity
	}

	tags = this.shouldTags(msg)
	if this.filter.TagsExclude {
		tags = !tags
	}
	//fmt.Printf("ShouldOutput: severity:%v, tags:%v, filter: %#v\n\n", severity, tags, this.filter)
	return tags
}

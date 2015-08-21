package sink

import (
	"github.com/chatoooo/logoon/core"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilterSeverity(t *testing.T) {
	assert := assert.New(t)

	filter := StandardSinkFilter{
		severityComparator: MakeStringSeverityComparator([]string{"TRACE", "DEBUG", "INFO", "WARNING", "ERROR", "FATAL"}),
		filter: core.Filter{
			Severity: core.FilterSeverity{
				CmpOp: core.FILTER_CMP_GE,
				Level: "INFO",
			},
			Tags:        nil,
			TagsExclude: false,
		},
	}
	assert.False(filter.ShouldOutput(core.CreateBasicLogMessage("TRACE", "")))
	assert.False(filter.ShouldOutput(core.CreateBasicLogMessage("DEBUG", "")))
	assert.True(filter.ShouldOutput(core.CreateBasicLogMessage("INFO", "")))
	assert.True(filter.ShouldOutput(core.CreateBasicLogMessage("WARNING", "")))
	assert.True(filter.ShouldOutput(core.CreateBasicLogMessage("ERROR", "")))
	assert.True(filter.ShouldOutput(core.CreateBasicLogMessage("FATAL", "")))
}

func TestFilterTags(t *testing.T) {
	assert := assert.New(t)

	filter := StandardSinkFilter{
		severityComparator: MakeStringSeverityComparator([]string{"TRACE", "DEBUG", "INFO", "WARNING", "ERROR", "FATAL"}),
		filter: core.Filter{
			Severity: core.FilterSeverity{
				CmpOp: core.FILTER_CMP_NONE,
				Level: "",
			},
			Tags:        []string{"ACCESS", "ERROR"},
			TagsExclude: false,
		},
	}
	assert.False(filter.ShouldOutput(core.CreateTaggedLogMessage("TRACE", "", []string{"SOMETHING"})))
	assert.False(filter.ShouldOutput(core.CreateTaggedLogMessage("DEBUG", "", []string{"SOMETHING"})))
	assert.True(filter.ShouldOutput(core.CreateTaggedLogMessage("INFO", "", []string{"ACCESS"})))
	assert.True(filter.ShouldOutput(core.CreateTaggedLogMessage("WARNING", "", []string{"ACCESS", "SOMETHING"})))
	assert.True(filter.ShouldOutput(core.CreateTaggedLogMessage("ERROR", "", []string{"SOMETHING", "ERROR"})))
	assert.False(filter.ShouldOutput(core.CreateTaggedLogMessage("FATAL", "", []string{"SOMETHING"})))
}

func TestFilterTagsExclude(t *testing.T) {
	assert := assert.New(t)

	filter := StandardSinkFilter{
		severityComparator: MakeStringSeverityComparator([]string{"TRACE", "DEBUG", "INFO", "WARNING", "ERROR", "FATAL"}),
		filter: core.Filter{
			Severity: core.FilterSeverity{
				CmpOp: core.FILTER_CMP_NONE,
				Level: "",
			},
			Tags:        []string{"ACCESS", "ERROR"},
			TagsExclude: true,
		},
	}
	assert.True(filter.ShouldOutput(core.CreateTaggedLogMessage("TRACE", "", []string{"SOMETHING"})))
	assert.True(filter.ShouldOutput(core.CreateTaggedLogMessage("DEBUG", "", []string{"SOMETHING"})))
	assert.False(filter.ShouldOutput(core.CreateTaggedLogMessage("INFO", "", []string{"ACCESS"})))
	assert.False(filter.ShouldOutput(core.CreateTaggedLogMessage("WARNING", "", []string{"ACCESS", "SOMETHING"})))
	assert.False(filter.ShouldOutput(core.CreateTaggedLogMessage("ERROR", "", []string{"SOMETHING", "ERROR"})))
	assert.True(filter.ShouldOutput(core.CreateTaggedLogMessage("FATAL", "", []string{"SOMETHING"})))
}

func TestFilterSeverityTags(t *testing.T) {
	assert := assert.New(t)

	filter := StandardSinkFilter{
		severityComparator: MakeStringSeverityComparator([]string{"TRACE", "DEBUG", "INFO", "WARNING", "ERROR", "FATAL"}),
		filter: core.Filter{
			Severity: core.FilterSeverity{
				CmpOp: core.FILTER_CMP_GT,
				Level: "INFO",
			},
			Tags:        []string{"ACCESS", "ERROR"},
			TagsExclude: false,
		},
	}
	assert.False(filter.ShouldOutput(core.CreateTaggedLogMessage("TRACE", "", []string{"SOMETHING", "ACCESS"})))
	assert.False(filter.ShouldOutput(core.CreateTaggedLogMessage("DEBUG", "", []string{"SOMETHING"})))
	assert.False(filter.ShouldOutput(core.CreateTaggedLogMessage("INFO", "", []string{"ACCESS"})))
	assert.True(filter.ShouldOutput(core.CreateTaggedLogMessage("WARNING", "", []string{"ACCESS", "SOMETHING"})))
	assert.True(filter.ShouldOutput(core.CreateTaggedLogMessage("ERROR", "", []string{"SOMETHING", "ERROR"})))
	assert.False(filter.ShouldOutput(core.CreateTaggedLogMessage("FATAL", "", []string{"SOMETHING"})))
}

func TestFilterSeverityTagsExclude(t *testing.T) {
	assert := assert.New(t)

	filter := StandardSinkFilter{
		severityComparator: MakeStringSeverityComparator([]string{"TRACE", "DEBUG", "INFO", "WARNING", "ERROR", "FATAL"}),
		filter: core.Filter{
			Severity: core.FilterSeverity{
				CmpOp: core.FILTER_CMP_GT,
				Level: "INFO",
			},
			Tags:        []string{"ACCESS", "ERROR"},
			TagsExclude: true,
		},
	}
	assert.False(filter.ShouldOutput(core.CreateTaggedLogMessage("TRACE", "", []string{"SOMETHING", "ACCESS"})))
	assert.False(filter.ShouldOutput(core.CreateTaggedLogMessage("DEBUG", "", []string{"SOMETHING"})))
	assert.False(filter.ShouldOutput(core.CreateTaggedLogMessage("INFO", "", []string{"ACCESS"})))
	assert.False(filter.ShouldOutput(core.CreateTaggedLogMessage("WARNING", "", []string{"ACCESS", "SOMETHING"})))
	assert.False(filter.ShouldOutput(core.CreateTaggedLogMessage("ERROR", "", []string{"SOMETHING", "ERROR"})))
	assert.True(filter.ShouldOutput(core.CreateTaggedLogMessage("FATAL", "", []string{"SOMETHING"})))
}

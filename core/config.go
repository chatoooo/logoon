package core

type SinkConfigFilter struct {
	Severity string   `json:"severity"`
	Tags     []string `json:"tags"`
	ExcludeTags bool `json:"exclude_tags"`
}

type SinkConfig struct {
	Disabled bool   `json:"disabled"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Filters  SinkConfigFilter `json:"filters"`
	ParsedFilter Filter      `json:"-"`
	Format       string      `json:"format"`
	Options      interface{} `json:"options"`
}

type Config struct {
	Severities []string     `json:"severities"`
	Debug bool `json:"debug"`
	Sinks      []*SinkConfig `json:"sinks"`
}

func ParseFilterConfig(config *SinkConfigFilter) Filter {
	var result Filter = Filter{}
	if config.Severity == ""{
		result.Severity.CmpOp = FILTER_CMP_NONE
	} else {
		firstChar := config.Severity[0]
		switch firstChar {
		case '>':
			result.Severity.CmpOp = FILTER_CMP_GT
			result.Severity.Level = config.Severity[1:]
		case '=':
			result.Severity.CmpOp = FILTER_CMP_EQ
			result.Severity.Level = config.Severity[1:]
		case '<':
			result.Severity.CmpOp = FILTER_CMP_LT
			result.Severity.Level = config.Severity[1:]
		default:
			result.Severity.CmpOp = FILTER_CMP_GE
			result.Severity.Level = config.Severity
		}
	}

	result.Tags = config.Tags
	result.TagsExclude = config.ExcludeTags
	return result
}

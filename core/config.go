package core

type SinkConfig struct {
	Disabled bool   `json:"disabled"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Filters  struct {
		Severity string   `json:"severity"`
		Tags     []string `json:"tags"`
	}
	ParsedFilter Filter `json:"-"`
	Format       string `json:"format"`
	Output       string `json:"output"`
}

type Config struct {
	Severities []string     `json:"severities"`
	Sinks      []SinkConfig `json:"sinks"`
}

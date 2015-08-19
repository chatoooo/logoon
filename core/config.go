package core

type SinkConfig struct {
	Disabled bool   `json:"disabled"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Filters  struct {
		Severity string   `json:"severity"`
		Tags     []string `json:"tags"`
	}
	Format  string      `json:"format"`
	Options interface{} `json:"options"`
}

type Config struct {
	Severities []string     `json:"severities"`
	Sinks      []SinkConfig `json:"sinks"`
}

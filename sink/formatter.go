package sink

import (
	"bytes"
	"github.com/chatoooo/logoon/core"
	"regexp"
	"strings"
	"text/template"
	"time"
)

type StandardSinkFormatter struct {
	compiledFormat   string
	compiledTemplate *template.Template
	timeFormat       string
}

func (this *StandardSinkFormatter) GetFormattedMessage(msg core.LogMessage) string {
	var b bytes.Buffer
	this.compiledTemplate.Execute(
		&b,
		map[string]interface{}{
			"time":     msg.Time().Format(this.timeFormat),
			"severity": msg.Severity(),
			"message":  strings.Replace(msg.Message(), "\n", "\n\t", -1),
			"tags":     strings.Join(msg.Tags(), ","),
		},
	)
	return b.String()
}

func (this *StandardSinkFormatter) SetFormat(format string) {
	timeFormatMatcher, _ := regexp.Compile("%time\\((.+?)\\)%")
	matcher, _ := regexp.Compile("%(.+?)(\\((.+?)\\))?%")

	timeFormat := timeFormatMatcher.FindStringSubmatch(format)
	if len(timeFormat) > 1 {
		this.timeFormat = timeFormat[1]
	} else {
		this.timeFormat = time.RFC3339
	}
	this.compiledFormat = matcher.ReplaceAllString(format, "{{.$1}}")
	this.compiledTemplate, _ = template.New("LogFormat").Parse(this.compiledFormat)
}

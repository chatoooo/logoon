package sink

import "github.com/chatoooo/logoon/core"

func containsTags(filter core.Filter, message core.LogMessage) bool {
	if filter.Tags != nil && len(filter.Tags) > 0 {
		for _, filterTag := range filter.Tags {
			for _, messageTag := range message.Tags() {
				if filterTag == messageTag {
					return true
				}
			}
		}
	}
	return false
}

func severityMatch(filter core.Filter, message core.LogMessage) {

}

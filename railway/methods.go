package railway

import (
	"encoding/json"

	"github.com/ferretcode/locomotive/util"
)

// searches for the given key and returns the corresponding value (and true) if found, or an empty string (and false)
func AttributesHasKeys(attributes []Attributes, keys []string) (string, bool) {
	for i := range attributes {
		for j := range keys {
			if keys[j] == attributes[i].Key {
				return attributes[i].Value, true
			}
		}
	}

	return "", false
}

func FilterLogs(logs []EnvironmentLog, wantedLevel []string, contentFilter string) []EnvironmentLog {
	if len(wantedLevel) == 0 && contentFilter == "" {
		return logs
	}

	filteredLogs := []EnvironmentLog{}

	for i := range logs {
		if !util.IsWantedLevel(wantedLevel, logs[i].Severity) {
			continue
		}

		// Convert log to JSON string for content filtering
		logJSON, _ := json.Marshal(logs[i])
		if !util.MatchesContentFilter(contentFilter, string(logJSON)) {
			continue
		}

		filteredLogs = append(filteredLogs, logs[i])
	}

	return filteredLogs
}

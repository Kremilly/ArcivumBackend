package utils

import (
	"encoding/json"
)

func Contains(rawMessage json.RawMessage, s string) bool {
	var emails []string
	if err := json.Unmarshal(rawMessage, &emails); err != nil {
		return false
	}

	for _, email := range emails {
		if email == s {
			return true
		}
	}

	return false
}
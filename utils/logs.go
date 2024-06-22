package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HeadersToString(headers http.Header) string {
	var result string

	if len(headers) == 0 {
		return result
	}

	for key, values := range headers {
		for _, value := range values {
			result += fmt.Sprintf("%s: %s\n", key, value)
		}
	}

	return result
}

func BodyToString(body interface{}) string {
	if body != nil {
		jsonBytes, _ := json.Marshal(body)
		return string(jsonBytes)
	}
	return ""
}

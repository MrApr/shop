package api

import (
	"shop/pkg/validation"
)

// generateResponse that is suitable for api and return it
func generateResponse(data any, message any) map[string]interface{} {
	return map[string]interface{}{
		"data":    data,
		"message": convertMessage(message),
	}
}

// convertMessage to string and return it
func convertMessage(message any) any {
	var convertedMsg string

	switch val := message.(type) {
	case error:
		convertedMsg = val.Error()
	case string:
		convertedMsg = val
	case []*validation.ValidationError:
		return val
	default:
		convertedMsg = ""
	}

	return convertedMsg
}

package tools

import (
	"encoding/json"
)

func ToMap(data any) map[string]any {
	result := map[string]any{}
	marshaled, _ := json.Marshal(data)
	json.Unmarshal(marshaled, &result)
	return result
}

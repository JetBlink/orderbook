package helper

import "encoding/json"

func Min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

// ToJsonString converts any value to JSON string.
func ToJsonString(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}

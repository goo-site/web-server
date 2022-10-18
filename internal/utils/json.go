package utils

import "encoding/json"

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func MarshalString(data interface{}) string {
	if b, err := json.Marshal(data); err == nil {
		return string(b)
	}
	return ""
}

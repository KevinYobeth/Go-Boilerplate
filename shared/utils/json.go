package utils

import "encoding/json"

func FromJsonString(jsonString string, target interface{}) error {
	return json.Unmarshal([]byte(jsonString), target)
}

func ToJsonString(target interface{}) (string, error) {
	jsonBytes, err := json.Marshal(target)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

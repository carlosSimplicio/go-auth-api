package utils

import "encoding/json"

func ParseJson(body []byte, dataStruct any) error {
	return json.Unmarshal(body, dataStruct)
}

package utils

import (
	"encoding/json"
	"log"
	"strconv"
)

// ConvertString to convert any data type to String
func ConvertString(v interface{}) string {
	result := ""
	if v == nil {
		return ""
	}
	switch v.(type) {
	case string:
		result = v.(string)
	case int:
		result = strconv.Itoa(v.(int))
	case int64:
		result = strconv.FormatInt(v.(int64), 10)
	case bool:
		result = strconv.FormatBool(v.(bool))
	case float64:
		result = strconv.FormatFloat(v.(float64), 'f', -1, 64)
	case []uint8:
		result = string(v.([]uint8))
	default:
		resultJSON, err := json.Marshal(v)
		if err == nil {
			result = string(resultJSON)
		} else {
			log.Println("Error on lib/converter ConvertString() ", err)
		}
	}

	return result
}

package mapper

import (
	"errors"
	"log"
	"strconv"
)

func getInt(obj map[string]interface{}, key string) (int, error) {
	switch t := obj[key].(type) {
	case string:
		val, err := strconv.Atoi(obj[key].(string))
		if err != nil {
			log.Printf("Error parsing int %s\n", obj[key].(string))
			return 0, err
		}
		return val, err
	case float64:
		val := int(obj[key].(float64))
		return val, nil
	default:
		log.Printf("Unknown field type for value %+v and type %s", obj[key], t)
		return 0, errors.New("Unable to parse")
	}
}

func getUInt(obj map[string]interface{}, key string) (uint32, error) {
	val, err := getInt(obj, key)
	if err != nil {
		return 0, err
	}
	return uint32(val), err
}

func getString(obj map[string]interface{}, key string) string {
	return obj[key].(string)
}

func getFloat(obj map[string]interface{}, key string) float64 {
	return obj[key].(float64)
}

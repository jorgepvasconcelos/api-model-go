package utils

import (
	"encoding/json"
	"log"
)

func SerializeToJson(data interface{}) []byte {
	serializedData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	return serializedData
}

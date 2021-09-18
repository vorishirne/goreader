package structures

import (
	"encoding/json"
	"log"
	"os"
)

func LoadJsonFile(fileName string, data interface{}) interface{} {
	fileText, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	if data == nil {
		data = map[string]interface{}{}
	}
	err = json.Unmarshal(fileText, data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

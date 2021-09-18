package structures

import (
	"encoding/json"
	"log"
	"os"
)

func LoadJsonFile(fileName string, data interface{}) interface{} {
	fileReader, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	if data == nil {
		data = &(map[string]interface{}{})
	}
	decoder := json.NewDecoder(fileReader)
	decoder.UseNumber()
	err = decoder.Decode(data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

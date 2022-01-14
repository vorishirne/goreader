package structure

import (
	"encoding/json"
	"log"
	"os"
)

/*
if a dataStructure is sent to "data", it will be loaded and
	nil will be returned
else,
	an a map is created and sent
*/

func LoadJsonFile(fileName string, data interface{}) (err error, inter interface{}) {
	fileReader, err := os.Open(fileName)
	if err != nil {
		return
	}
	var dataWasNil bool
	if data == nil {
		dataWasNil = true
		var input interface{}
		data = &input
	}
	decoder := json.NewDecoder(fileReader)
	decoder.UseNumber()
	err = decoder.Decode(data)
	if err != nil {
		return
	}
	if dataWasNil {
		return nil, data
	}
	return nil, nil
}

func WriteJson(fileName string, data interface{}) {
	yamlBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(fileName, yamlBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}

}

package structures

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

func WriteYaml(fileName string, data interface{}) {
	yamlBytes, err := yaml.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(fileName, yamlBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}

}

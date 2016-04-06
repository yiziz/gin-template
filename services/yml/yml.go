package yml

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// ReadConfigFile reads in a file of name
func ReadConfigFile(name string) []byte {
	file, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

// ConfigYML returns a map[interface{}]interface{} of the yaml file with name
func ConfigYML(name string) map[interface{}]interface{} {
	m := make(map[interface{}]interface{})
	err := yaml.Unmarshal(ReadConfigFile(name), &m)
	if err != nil {
		log.Fatal(err)
	}
	return m
}

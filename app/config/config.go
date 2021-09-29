package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ConfigService struct {
	Character  []string `json:"character"`
	Eyewear    []string `json:"eyewear"`
	FrontItem  []string `json:"frontitem"`
	Hands      []string `json:"hands"`
	Background []string `json:"background"`
}

func NewConfigService(configPath string) *ConfigService {
	jsonFile, err := os.Open(configPath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var service ConfigService

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &service)

	return &service
}

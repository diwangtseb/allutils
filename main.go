package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/diwangtseb/yamldemo/conf"
	yaml "gopkg.in/yaml.v2"
)

var confYamlModel = conf.ConfYamlModel{
	WhiteList: []string{},
}

func main() {
	file, err := ioutil.ReadFile("./conf/app.yaml")
	if err != nil {
		log.Fatal("fail to read file:", err)
	}
	err = yaml.Unmarshal(file, &confYamlModel)
	if err != nil {
		log.Fatal("fail to unmarshal file:", err)
	}
	fmt.Printf("%v\n", confYamlModel.WhiteList)
	fmt.Println(!Contains(confYamlModel.WhiteList, "18627079270"))
}

func Contains(T interface{}, checkObj interface{}) bool {
	switch T.(type) {
	case []string:
		switch checkObj.(type) {
		case string:
			for _, value := range T.([]string) {
				if value != checkObj {
					continue
				}
				return true
			}
			return false
		}
	}
	return false
}

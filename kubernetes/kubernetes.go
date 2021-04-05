package main

import (
	"io/ioutil"

	"github.com/ghodss/yaml"
)

func main() {
	bytes, err := ioutil.ReadFile("deployment.yml")
	if err != nil {
		panic(err.Error())
	}

	var spec v1.Deployment
	err = yaml.Unmarshal(bytes, &spec)
	if err != nil {
		panic(err.Error())
	}
}

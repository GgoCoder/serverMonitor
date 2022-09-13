package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v3"
)

func ReadConfig()*ConfigYaml{
	yamlFile, err := ioutil.ReadFile("./../config.yaml")
	if err != nil {
		return nil
	}

	var config *ConfigYaml

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil
	}
	return config
}

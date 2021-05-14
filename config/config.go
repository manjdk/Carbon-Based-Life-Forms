package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	DatabaseHost string `yaml:"database_host" validate:"required"`
	AwsRegion    string `yaml:"region" validate:"required"`
	QueueFactory Queue  `yaml:"queue_factory" validate:"required"`
	QueueManager Queue  `yaml:"queue_manager" validate:"required"`
	MineralTable string `yaml:"mineral_table" validate:"required"`
	ManagerURL   string `yaml:"manager_client_url" validate:"required"`
	FactoryURL   string `yaml:"factory_client_url" validate:"required"`
}

type Queue struct {
	Name string `yaml:"name" validate:"required"`
	Host string `yaml:"host" validate:"required"`
}

func NewConfig(path string) (*Config, error) {
	source, err := ioutil.ReadFile(path + "/config.yaml")
	if err != nil {
		return nil, err
	}

	c := new(Config)
	return c, yaml.Unmarshal(source, c)
}

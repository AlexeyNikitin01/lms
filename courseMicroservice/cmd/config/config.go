package config

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

const (
	ConfigPath string = "./etc/config.template.yml"
)

type CourseMicroservice struct {
	Postgres *Postgres `yaml:"psql"`
	Mongo    *Mongo    `yaml:"mongo"`
}

type Postgres struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	DBName   string `yaml:"dbname"`
	Password string `yaml:"pass"`
	SSLmode  string `yaml:"sslmode"`
}

type Mongo struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"pass"`
}

// NewConfigCourseMicroservice
//
// TODO: необходимо переместить ключи доступа в окружение github. Ключи видны всем. Это не правильно.
func NewConfigCourseMicroservice() (*CourseMicroservice, error) {
	yamlFile, err := os.ReadFile(ConfigPath)
	if err != nil {
		return nil, errors.Wrap(err, "read file config")
	}

	var config *CourseMicroservice

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal")
	}

	return config, nil
}

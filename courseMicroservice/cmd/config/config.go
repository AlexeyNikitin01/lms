package config

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

var configs = []string{
	"./etc/config.template.yml",
	"./etc/config.aws.yml",
}

type CourseMicroservice struct {
	Postgres *Postgres `yaml:"psql"`
	Mongo    *Mongo    `yaml:"mongo"`
	AWS      *AWS      `yaml:"AWS"`
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

type AWS struct {
	PublicKey string `yaml:"publickey"`
	SecretKey string `yaml:"secretkey"`
	Region    string `yaml:"region"`
	Endpoint  string `yaml:"endpoint"`
	Bucket    string `yaml:"bucket"`
}

// NewConfigCourseMicroservice
//
// TODO: необходимо переместить ключи доступа в окружение github. Ключи видны всем. Это не правильно.
func NewConfigCourseMicroservice() (*CourseMicroservice, error) {
	var courseMicroservice *CourseMicroservice

	for _, config := range configs {
		yamlFile, err := os.ReadFile(config)
		if err != nil {
			return nil, errors.Wrap(err, "read file config")
		}

		err = yaml.Unmarshal(yamlFile, &courseMicroservice)
		if err != nil {
			return nil, errors.Wrap(err, "unmarshal")
		}
	}

	return courseMicroservice, nil
}

package config

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

var configs = []string{
	"./etc/config.template.yml",
}

const (
	PathPostgres string = "./etc/config.template.yml"
)

type LabMic struct {
	Postgres *Postgres `yaml:"psql"`
}

type Postgres struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	DBName   string `yaml:"dbname"`
	Password string `yaml:"pass"`
	SSLmode  string `yaml:"sslmode"`
}

// NewLabCourseMicroservice
//
// TODO: необходимо переместить ключи доступа в окружение github. Ключи видны всем. Это не правильно.
func NewLabCourseMicroservice() (*LabMic, error) {
	var courseMicroservice *LabMic

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

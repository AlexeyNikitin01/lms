package config

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

const (
	ConfigPath string = "./etc/config.template.yml"
)

type UserMicroservice struct {
	Postgres *Postgres `yaml:"psql"`
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

type AWS struct {
	PublicKey string `yaml:"publickey"`
	SecretKey string `yaml:"secretkey"`
	Region    string `yaml:"region"`
	Endpoint  string `yaml:"endpoint"`
	Bucket    string `yaml:"bucket"`
}

// NewConfigUserMicroservice
//
// TODO: необходимо переместить ключи доступа в окружение github. Ключи видны всем. Это не правильно.
func NewConfigUserMicroservice() (*UserMicroservice, error) {
	yamlFile, err := os.ReadFile(ConfigPath)
	if err != nil {
		return nil, errors.Wrap(err, "read file config")
	}

	var config *UserMicroservice

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal")
	}

	return config, nil
}

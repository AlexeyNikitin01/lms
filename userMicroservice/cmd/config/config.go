package config

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

const (
	PathPostgres string = "./etc/config.template.yml"
	PathAWS      string = "./etc/config.aws.yml"
)

type UserMicroservicePostgres struct {
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

// NewCfgPostgres
//
// TODO: необходимо переместить ключи доступа в окружение github. Ключи видны всем. Это не правильно.
func NewCfgPostgres() (*UserMicroservicePostgres, error) {
	yamlFile, err := os.ReadFile(PathPostgres)
	if err != nil {
		return nil, errors.Wrap(err, "read file config")
	}

	var config *UserMicroservicePostgres

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal")
	}

	return config, nil
}

type UserMicroserviceAWS struct {
	AWS *AWS `yaml:"AWS"`
}

type AWS struct {
	PublicKey string `yaml:"publickey"`
	SecretKey string `yaml:"secretkey"`
	Region    string `yaml:"region"`
	Endpoint  string `yaml:"endpoint"`
	Bucket    string `yaml:"bucket"`
	Active    bool
}

// NewCfgAWS если нет конфигурации, то файлы сохраняются локально.
func NewCfgAWS() (*UserMicroserviceAWS, error) {
	yamlFile, err := os.ReadFile(PathAWS)
	if errors.Is(err, os.ErrNotExist) {
		return &UserMicroserviceAWS{&AWS{}}, nil
	} else if err != nil {
		return nil, errors.Wrap(err, "read file config")
	}

	var config *UserMicroserviceAWS

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal")
	}

	config.AWS.Active = true

	return config, nil
}

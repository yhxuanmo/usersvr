package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Conf struct {
	Server struct {
		Http	string 	`yaml:"http"`
		Grpc 	string 	`yaml:"grpc"`
		Debug   bool 	`yaml:"debug"`
	}

	Postgresql struct{
		Host		string	`yalm:"host"`
		Port		string	`yalm:"port"`
		User		string	`yalm:"user"`
		DbName		string	`yalm:"dbname"`
		Password	string	`yalm:"password"`
		Sslmode		string	`yalm:"sslmode"`
	}

	Redis struct{
		Addr		string	`yaml:"addr"`
		Password	string	`yaml:"password"`
		Db		int	`yaml:"db"`
	}

	Email struct{
		User 	string	`yaml:"user"`
		Pass 	string	`yaml:"pass"`
		Host 	string	`yaml:"host"`
		Port 	string	`yaml:"port"`
	}

	Jwt struct{
		Secret 	string	`yaml:"secret"`
	}
}

var Config Conf

func init() {
	yamlFile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		log.Fatalf("config get err %s", err)
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		log.Fatalf("Unmarshal err %s", err)
	}
}

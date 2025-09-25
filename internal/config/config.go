package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	Env 		string			`env:"ENV" evDefault:"dev"`
	Host    	string	        `env:"HOST" envDefault:"localhost"`
	Port		string			`env:"PORT" envDefault:"8080"`
	Mysql     	MysqlConfig		`envPrefix:"mysql_"`
	JWTSecretKey string 		`env:"JWT_SECRET_KEY" envDefault:"jwt"`
}

type MysqlConfig struct {
	Host 		string 			`env:"HOST" envDefault:"localhost"`
	Port		string			`env:"PORT" envDefault:"3306"`
	User		string			`env:"USER" envDefault:"root"`
	Password	string			`env:"PASSWORD" envDefault:""`
	Database	string			`env:"DATABASE" envDefault:"evermos_rakamin"`
}

func NewConfig(envPath string) (*Config, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		return nil, err
	}

	cfg := new(Config)
	err = env.Parse(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
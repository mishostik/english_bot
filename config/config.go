package config

import (
	"english_bot/pkg/secure"
	"github.com/spf13/viper"
	"log"
)

// TODO MONGO config
type Config struct {
	Postgres PostgresConfig `mapstructure:"postgres"`
	Telegram TelegramConfig `mapstructure:"telegram"`
	Logger   LoggerConfig   `mapstructure:"logger"`
	Mongo    MongoConfig    `mapstructure:"mongo"`
}

type MongoConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"-"`
	DBName   string `json:"DBName"`
}

type LoggerConfig struct {
	Level             string `yaml:"level"`
	InFile            string `yaml:"inFile"`
	Development       bool   `yaml:"development"`
	DisableCaller     bool   `yaml:"disableCaller"`
	DisableStacktrace bool   `yaml:"disableStacktrace"`
	Encoding          string `yaml:"encoding"`
}

type PostgresConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"-"`
	DBName   string `json:"DBName"`
	SSLMode  string `json:"sslMode"`
	PgDriver string `json:"pgDriver"`
}

type TelegramConfig struct {
	Token string `json:"token"`
}

func LoadConfig() (*viper.Viper, error) {

	viperInstance := viper.New()

	viperInstance.AddConfigPath("./config")
	viperInstance.SetConfigName("config")
	viperInstance.SetConfigType("yml")

	err := viperInstance.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return viperInstance, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {

	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Fatalf("unable to decode config into struct, %v", err)
		return nil, err
	}
	return &c, nil
}

func DecryptConfig(cfg *Config, s *secure.Shield) {

	cfg.Postgres.Host = s.DecryptMessage(cfg.Postgres.Host)
	cfg.Postgres.Port = s.DecryptMessage(cfg.Postgres.Port)
	cfg.Postgres.User = s.DecryptMessage(cfg.Postgres.User)
	cfg.Postgres.Password = s.DecryptMessage(cfg.Postgres.Password)
	cfg.Postgres.DBName = s.DecryptMessage(cfg.Postgres.DBName)

	cfg.Telegram.Token = s.DecryptMessage(cfg.Telegram.Token)
}

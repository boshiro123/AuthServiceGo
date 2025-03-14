package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

var configFlags = flag.NewFlagSet("config", flag.ContinueOnError)
var configPath string

func init() {
	configFlags.StringVar(&configPath, "config", "", "path to config file")
}

type Config struct {
	Env      string        `yaml:"env" env-default:"local"`
	TokenTTL time.Duration `yaml:"token_ttl"`
	Secrets  Secrets       `yaml:"secrets"`
	Postgres Postgres      `yaml:"postgres"`
	App      App           `yaml:"app"`
}

type Secrets struct {
	SecretKey string `yaml:"jwt"`
	Salt      string `yaml:"salt"`
}

type Postgres struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

type App struct {
	Port string `yaml:"port"`
}

func MustLoad() *Config {
	path := getConfigPath()

	if path == "" {
		panic("config file path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist" + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config" + err.Error())
	}

	return &cfg
}

func getConfigPath() string {
	return os.Getenv("CONFIG_PATH")
}

package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Env string `yaml: "env" env-default:"local"`
	HTTPServer `yaml: "http_server"`
}


type HTTPServer struct {
	Address string `yaml:"address" env-default:"localhost:8083"`
	Port string `yaml:"port" env-default:"8083"`

	Timeout time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MastLoad() *Config{

	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	// Check for path
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == ""{
		log.Fatalf("config path not set")
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err){
		log.Fatalf("config file doesn't exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil{
		log.Fatalf("can't read config : %s", err)
	}

	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	err = viper.ReadInConfig()
	if err != nil{
		log.Fatal("can't read viper config: %v", err)
	}

	return &cfg
}
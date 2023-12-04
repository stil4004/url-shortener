package main

import (
	"log/slog"
	"os"

	"github.com/spf13/viper"
	"github.com/stil4004/url-shorter/internal/config"
	"github.com/stil4004/url-shorter/internal/db"
)

const(
	envLocal = "local"
	envDev = "dev"
)

func main() {
	cfg := config.MastLoad()
	//fmt.Println(cfg)
	log := setupLogger(cfg.Env)

	log.Info("starting app")
	log.Debug("started debug messages")

	dataBase, err := db.New(db.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password:  os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Error("failed to init storage ", err)
		os.Exit(1)
	}

	_ = dataBase
}

func setupLogger(env string) *slog.Logger { 
	var log *slog.Logger

	switch env{
	case envLocal :
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}
	return log
}
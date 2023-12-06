package main

import (
	"database/sql"
	"log/slog"
	"os"

	"github.com/spf13/viper"
	urlshorter "github.com/stil4004/url-shorter"
	"github.com/stil4004/url-shorter/internal/config"
	"github.com/stil4004/url-shorter/internal/handler"
	"github.com/stil4004/url-shorter/internal/repository"
	"github.com/stil4004/url-shorter/internal/repository/db"
	"github.com/stil4004/url-shorter/internal/service"
)

const(
	envLocal = "local"
	envDev = "dev"
)

func main() {

	memory_arg := checkMemoryArgs(os.Args[1:])
	
	cfg := config.MastLoad()
	//fmt.Println(cfg)
	log := setupLogger(cfg.Env)

	log.Info("starting app")
	log.Debug("started debug messages")

	var data_base *sql.DB
	if memory_arg == "db" {
		var err error
		data_base, err = db.New(db.Config{
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
	} 

	// Change data-base to Abstract structure DB
	repos := repository.NewRepository(data_base, memory_arg)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(urlshorter.Server)
	err := srv.Run("8082", handlers.InitRoutes())
	if err != nil{
		log.Error("couldn't run server: %v", err)
		os.Exit(1)
	}
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

func checkMemoryArgs(args []string) string{	
	for _, arr := range args{
		if arr == "-im" || arr == "-in-memory"{
			return "im"
		}
		if arr == "-db" || arr == "-data-base"{
			return "db"
		}
	}
	return "db"
}
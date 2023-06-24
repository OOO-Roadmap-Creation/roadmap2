package main

import (
	"context"
	"github.com/OOO-Roadmap-Creation/roadmap2"
	"github.com/OOO-Roadmap-Creation/roadmap2/consumer"
	"github.com/OOO-Roadmap-Creation/roadmap2/pkg/handler"
	"github.com/OOO-Roadmap-Creation/roadmap2/pkg/repository"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title Roadmap2 API
// @version 1.0
// @description Публичные роадмапы, ноды, рейтинг

// @host localhost:8081
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		//Password: os.Getenv("DB_PASSWORD"),
		Password: "qwerty",
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	handlers := handler.NewHandler(repos)
	go consumer.Recoverer(10, 1, consumer.Listen, repos)

	srv := new(roadmap2.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("TodoApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("TodoApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"leadgen/pkg/config"
	"leadgen/pkg/db"
	"leadgen/pkg/handler"
	"leadgen/pkg/repository"
	"leadgen/pkg/service"
	"log"
	"os"
)

type App struct {
	cfg *viper.Viper
	log *logrus.Logger

	server *handler.ApiHandler
}

func NewApp() *App {
	cfg := config.New()
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetReportCaller(true)
	log.SetLevel(logrus.InfoLevel)

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db := db.NewPgDb(
		log,
		cfg.GetString("database.url"),
		cfg.GetString("database.user"),
		cfg.GetString("database.pass"),
		cfg.GetString("database.database"),
	)

	r := repository.New(log, db)
	s := service.New(log, r)
	g := gin.New()

	server := handler.NewApiHandler(g, log, s, 8080)

	return &App{
		cfg:    cfg,
		log:    log,
		server: server,
	}
}

func (a *App) Run(sigterm chan os.Signal) {
	a.server.Run()

	<-sigterm

	err := a.server.Shutdown()
	if err != nil {
		a.log.WithError(err).Error("shutdown error")
	}
	log.Println("terminating: via signal")
}

package main

import (
	"context"
	"io"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	casperlib "github.com/Simplewallethq/simple-wallet-tg-bot/library/client"
	"github.com/Simplewallethq/simple-wallet-tg-bot/rest-api/blockchain/casper"
	swagdocsfile "github.com/Simplewallethq/simple-wallet-tg-bot/rest-api/docs"
	"github.com/Simplewallethq/simple-wallet-tg-bot/rest-api/handlers"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title CSPR REST
// @version 1.0
// @description casper blockchain rest api

// @host localhost:8081
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-Token

func main() {
	logger := NewLogger()

	casperclient := casperlib.NewClient()

	casperBlockchain := casper.New(casperclient, logger)
	swagdoc := ginSwagger.WrapHandler(swaggerFiles.Handler)
	swagdocsfile.SwaggerInfo.Host = ""
	swagdocsfile.SwaggerInfo.BasePath = "/"

	handler := handlers.NewHandler(logger, casperBlockchain, casperBlockchain, swagdoc, GetAuthConf())

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Create a context to propagate the shutdown signal
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-quit
		cancel()
	}()

	config := getGinConfig()

	err := handler.Start(ctx, config)
	if err != nil {
		logger.Fatal(err)
	}

}

func NewLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{})

	file, err := os.OpenFile("log.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		logger.SetOutput(file)
	} else {
		logger.Info("Failed to log to file, using default stderr")
	}

	logger.SetOutput(io.MultiWriter(os.Stdout, file))
	return logger
}

func getGinConfig() handlers.GinConfig {
	port := os.Getenv("REST_PORT")
	if port == "" {
		port = ":8081"
	}

	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = gin.DebugMode
	}

	return handlers.GinConfig{
		Port: port,
		Mode: mode,
	}
}

func GetAuthConf() handlers.Auth {
	needAuthStr := os.Getenv("ENABLE_AUTH")
	needAuth := false
	if needAuthStr != "" {
		needAuth, _ = strconv.ParseBool(needAuthStr)
	}
	Token := os.Getenv("TOKEN")
	//log.Println("needAuth:", needAuth, "Token:", Token)
	return handlers.Auth{
		Auth:  needAuth,
		Token: Token,
	}
}

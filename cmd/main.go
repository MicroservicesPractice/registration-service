package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"

	"registration-service/api"
	"registration-service/cmd/config"
	"registration-service/shared/helpers"
)

var SERVER_PORT = helpers.GetEnv("SERVER_PORT")

func init() {
	helpers.CheckRequiredEnvs()

	config.InitLogger()
}

func main() {
	dataBase := config.ConnectDb()

	defer dataBase.Close()

	router := gin.Default()

	api.Handlers(router, dataBase)

	err := router.Run(fmt.Sprintf(":%v", SERVER_PORT))

	if err != nil {
		log.Panicf("Server listen err: %v", err)
	}

	log.Infof("Server has been started on port %v", SERVER_PORT)
}

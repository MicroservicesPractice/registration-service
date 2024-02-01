package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"

	"registration-service/app/api"
	"registration-service/app/config/initializers"
	"registration-service/app/helpers"
)

var SERVER_PORT = helpers.GetEnv("SERVER_PORT")

func init() {
	helpers.CheckRequiredEnvs()

	initializers.InitLogger()
}

func main() {
	dataBase := initializers.ConnectDb()

	defer dataBase.Close()

	router := gin.Default()

	log.Infof("FUUUUUUCK")

	api.Controllers(router, dataBase)

	err := router.Run(fmt.Sprintf(":%v", SERVER_PORT))

	if err != nil {
		log.Panicf("Server listen err: %v", err)
	}

	log.Infof("Server has been started on port %v", SERVER_PORT)
}

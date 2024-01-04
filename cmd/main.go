package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"

	"exampleApi/api"
	"exampleApi/cmd/config"
	"exampleApi/shared/helpers"
)

var SERVER_PORT = helpers.GetEnv("SERVER_PORT")

func init() {
	helpers.CheckRequiredEnvs()

	config.InitLogger()
}

func main() {
	redisDb := config.ConnectRedis()

	dataBase := config.ConnectDb()

	defer dataBase.Close()
	defer redisDb.Close()

	router := gin.Default()

	api.Handlers(router, dataBase, redisDb)

	err := router.Run(fmt.Sprintf(":%v", SERVER_PORT))

	if err != nil {
		log.Panicf("Server listen err: %v", err)
	}

	log.Infof("Server has been started on port %v", SERVER_PORT)
}

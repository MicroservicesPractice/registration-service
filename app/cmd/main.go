package main

import (
	"context"
	"registration-service/app/proto"

	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"

	"registration-service/app/config/initializers"
	"registration-service/app/helpers"
)

var SERVER_PORT = helpers.GetEnv("SERVER_PORT")

func init() {
	helpers.CheckRequiredEnvs()

	initializers.InitLogger()
}

func sendGrpcRequ() {
	conn, err := grpc.Dial("localhost:5003", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewUserClient(conn)

	createUserReq := &proto.CreateUserRequest{
		Email:       "user@example.com",
		Password:    "password",
		Nickname:    "nickname",
		PhoneNumber: "123456789",
	}
	createUserRes, err := client.CreateUser(context.Background(), createUserReq)
	if err != nil {
		log.Fatalf("CreateUser request failed: %v", err)
	}
	log.Printf("CreateUser response: %s", createUserRes.Message)

	// Example: Call GetUserPassword
	getUserPasswordReq := &proto.GetUserPasswordRequest{Id: "user_id"}
	getUserPasswordRes, err := client.GetUserPassword(context.Background(), getUserPasswordReq)
	if err != nil {
		log.Fatalf("GetUserPassword request failed: %v", err)
	}
	log.Printf("GetUserPassword response: %s", getUserPasswordRes.Password)
}

func main() {
	// dataBase := initializers.ConnectDb()

	// defer dataBase.Close()

	sendGrpcRequ()

	// router := gin.Default()

	log.Infof("FUUUUUUCK")

	// api.Controllers(router, dataBase)

	// err := router.Run(fmt.Sprintf(":%v", SERVER_PORT))

	// if err != nil {
	// 	log.Panicf("Server listen err: %v", err)
	// }

	// log.Infof("Server has been started on port %v", SERVER_PORT)
}

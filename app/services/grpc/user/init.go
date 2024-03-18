package user

import (
	"fmt"
	"registration-service/app/helpers"
	"registration-service/app/helpers/log"

	"google.golang.org/grpc"
)

var USER_SERVICE_HOST = helpers.GetEnv("USER_SERVICE_HOST")
var USER_SERVICE_GRPC_PORT = helpers.GetEnv("USER_SERVICE_GRPC_PORT")

func ConnectUserServiceGrpc() UserClient {
	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", USER_SERVICE_HOST, USER_SERVICE_GRPC_PORT), grpc.WithInsecure())
	if err != nil {
		log.GrpcLog(log.Error, "user-service", "can't connect to grpc service")
	}
	// defer conn.Close()

	client := NewUserClient(conn)

	return client

	// Example: Call GetUserPassword
	// getUserPasswordReq := &proto.GetUserPasswordRequest{Id: "user_id"}
	// getUserPasswordRes, err := client.GetUserPassword(context.Background(), getUserPasswordReq)
	// if err != nil {
	// 	log.Fatalf("GetUserPassword request failed: %v", err)
	// }
	// log.Printf("GetUserPassword response: %s", getUserPasswordRes.Password)
}

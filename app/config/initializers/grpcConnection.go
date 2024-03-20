package initializers

import (
	"fmt"
	"registration-service/app/helpers"
	"registration-service/app/helpers/log"

	grpcApi "github.com/MicroservicesPractice/grpc-api/generated/user"

	"google.golang.org/grpc"
)

var USER_SERVICE_HOST = helpers.GetEnv("USER_SERVICE_HOST")
var USER_SERVICE_GRPC_PORT = helpers.GetEnv("USER_SERVICE_GRPC_PORT")

func ConnectUserServiceGrpc() grpcApi.UserClient {
	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", USER_SERVICE_HOST, USER_SERVICE_GRPC_PORT), grpc.WithInsecure())
	if err != nil {
		log.GrpcLog(log.Error, "user-service", "can't connect to grpc service")
	}
	// defer conn.Close()

	client := grpcApi.NewUserClient(conn)

	return client
}

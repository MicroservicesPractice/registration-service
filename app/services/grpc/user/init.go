package user

import (
	"registration-service/app/helpers/log"

	"google.golang.org/grpc"
)

func ConnectUserServiceGrpc() UserClient {
	conn, err := grpc.Dial("localhost:6003", grpc.WithInsecure())
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

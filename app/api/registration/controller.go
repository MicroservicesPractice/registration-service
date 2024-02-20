package registration

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"registration-service/app/consts"
	"registration-service/app/helpers"
	"registration-service/app/helpers/log"
	userGrpc "registration-service/app/services/grpc/user"
)

func SignUp(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": consts.INVALID_BODY})
		log.HttpLog(c, log.Warn, http.StatusBadRequest, fmt.Sprintf("%v: %v", consts.INVALID_BODY, err.Error()))
		return
	}

	validationResult := helpers.Validate(&user)
	if !validationResult.OK {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationResult.Errors})
		log.HttpLog(c, log.Warn, http.StatusBadRequest, "validation error")
		return
	}

	password, err := helpers.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": consts.SOMETHING_WENT_WRONG})
		log.HttpLog(c, log.Error, http.StatusInternalServerError, fmt.Sprintf("can't hash password: %v", err.Error()))
		return
	}

	user.Password = password

	grpcClient := userGrpc.ConnectUserServiceGrpc()

	createUserReq := &userGrpc.CreateUserRequest{
		Email:       user.Email,
		Password:    user.Password,
		Nickname:    user.Nickname,
		PhoneNumber: user.PhoneNumber,
	}

	createUserRes, err := grpcClient.CreateUser(context.Background(), createUserReq)
	if err != nil {
		log.GrpcLog(log.Error, "user-service", fmt.Sprintf("CreateUser request failed: %v", err))
	}
	log.GrpcLog(log.Info, "user-service", fmt.Sprintf("CreateUser response: %s", createUserRes.Message))

	c.JSON(http.StatusOK, gin.H{"message": "user was created successfully"})
	log.HttpLog(c, log.Info, http.StatusOK, fmt.Sprintf("user was created successfully: %v", user.Email))
}

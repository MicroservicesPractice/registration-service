package registration

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	grpcApi "github.com/MicroservicesPractice/grpc-api/generated/user"

	"github.com/gin-gonic/gin"

	"registration-service/app/config/initializers"
	"registration-service/app/consts"
	"registration-service/app/helpers"
	"registration-service/app/helpers/log"

	amqp "github.com/rabbitmq/amqp091-go"
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

	grpcClient := initializers.ConnectUserServiceGrpc()

	createUserReq := &grpcApi.CreateUserRequest{
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

func SendRmqMessage(c *gin.Context) {
	type Message struct {
		Message string `json:"message"`
	}
	var data Message

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": consts.INVALID_BODY})
		log.HttpLog(c, log.Warn, http.StatusBadRequest, fmt.Sprintf("%v: %v", consts.INVALID_BODY, err.Error()))
		return
	}

	// connect
	conn, err := amqp.Dial("amqp://root:1234@localhost:5672/") // Создаем подключение к RabbitMQ
	if err != nil {
		fmt.Printf("unable to open connect to RabbitMQ server. Error: %s", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("failed to open channel. Error: %s", err)
	}

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// msg := Message{
	// 	Message: "Hello, RabbitMQ!",
	// }

	// jsonMsg, err := json.Marshal(msg)
	// if err != nil {
	// 	fmt.Printf("Failed to marshal JSON: %v", err)
	// }

	// body := "Hello World!"
	// err = ch.PublishWithContext(ctx,
	// 	"notification",     // exchange
	// 	"confirm.password", // routing key
	// 	false,              // mandatory
	// 	false,              // immediate
	// 	amqp.Publishing{
	// 		ContentType: "application/json",
	// 		Body:        jsonMsg,
	// 	})
	// if err != nil {
	// 	fmt.Printf("failed to publish a message. Error: %s", err)
	// }

	var target = map[string]string{
		"EMAIL_CONFIRMATION_COMPLETE": "EMAIL_CONFIRMATION_COMPLETE-success",
	}
	jsonMsg2, err := json.Marshal(target)
	fmt.Println("hhe", jsonMsg2)
	if err != nil {
		fmt.Printf("Failed to marshal JSON: %v", err)
	}

	ctx2, cancel2 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel2()

	err = ch.PublishWithContext(ctx2,
		"notification",      // exchange
		"email.status.info", // routing key
		false,               // mandatory
		false,               // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonMsg2,
		})
	if err != nil {
		fmt.Printf("failed to publish a message. Error: %s", err)
	}

	defer func() {
		_ = ch.Close() // Закрываем канал в случае удачной попытки открытия
	}()

	defer func() {
		_ = conn.Close() // Закрываем подключение в случае удачной попытки
	}()

	c.JSON(http.StatusOK, gin.H{"message": "user was created successfully"})
}

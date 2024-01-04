package helpers

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"

	log "github.com/sirupsen/logrus"
)

func GetEnv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Panicf("Error loading .env file: %v", err)
	}

	return os.Getenv(key)
}

func CheckRequiredEnvs() {
	requiredEnvVars := []string{"SERVER_PORT", "DB_PORT", "DB_HOST", "DB_NAME", "DB_USER", "DB_PASSWORD", "LOG_LEVEL", "REDIS_HOST", "REDIS_PORT", "REDIS_PASSWORD", "ACCESS_TOKEN_EXP_MIN", "REFRESH_TOKEN_SECRET", "REFRESH_TOKEN_EXP_MIN", "ACCESS_TOKEN_PUBLIC_SECRET_PATH", "ACCESS_TOKEN_PRIVATE_SECRET_PATH"}

	for _, envVar := range requiredEnvVars {
		if value, exists := os.LookupEnv(envVar); !exists || value == "" {
			log.Panic(fmt.Sprintf("Error: Environment variable %v is not set.", envVar))
		}
	}
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashedPassword), err
}

func CheckPassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}

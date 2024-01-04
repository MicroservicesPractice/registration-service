package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"registration-service/api/registration"
)

func Handlers(r *gin.Engine, db *sql.DB) {

	InitMiddlewares(r, db)

	r.POST("/registration/signUp", registration.SignUp)
}

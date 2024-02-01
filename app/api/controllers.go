package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"registration-service/app/api/registration"
)

func Controllers(r *gin.Engine, db *sql.DB) {

	InitMiddlewares(r, db)

	r.POST("/registration/signUp", registration.SignUp)
}

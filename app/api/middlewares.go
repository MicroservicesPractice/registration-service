package api

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
)

func InitMiddlewares(r *gin.Engine, db *sql.DB) {
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// set start time on every request to check response time
	r.Use(func(c *gin.Context) {
		startTime := time.Now()
		c.Set("startTime", startTime)
		c.Next()
	})
}

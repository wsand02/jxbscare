package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wsand02/jxbscare/web/handlers"
)

func InitRoutes(r *gin.Engine) {
	r.GET("/ping", handlers.Ping)
}

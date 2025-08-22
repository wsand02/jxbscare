package web

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Web struct {
	router     *gin.Engine
	routesInit bool
}

func InitWeb() *Web {
	router := gin.Default()
	w := Web{router: router}
	w.routesInit = false
	return &w
}

func (w *Web) InitRoutes() {
	w.router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	w.routesInit = true
}

func (w *Web) StartServer() {
	if !w.routesInit {
		log.Fatal("Routes uninitialized!")
		return
	}
	err := w.router.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		log.Fatal(err)
	}
}

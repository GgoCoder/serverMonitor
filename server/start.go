package server

import (
	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	r.POST("/login", Login)
	r.GET("/home", ListService)

	r.GET("/service", ListService)
	r.POST("/service", AddService)
	r.PUT("/service", UpdateService)
	r.DELETE("/service", DeleteService)

	r.Run(":8000")
}

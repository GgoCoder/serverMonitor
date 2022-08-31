package main

import (
	"serverMonitor/server"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/login", server.Login)
	r.GET("/home", server.ListService)

	r.GET("/service", server.ListService)
	r.POST("/service", server.AddService)
	r.PUT("/service", server.UpdateService)
	r.DELETE("/service", server.DeleteService)

	r.Run(":8000")
}

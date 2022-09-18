package router

import (
	"github.com/gin-gonic/gin"
	"serverMonitor/internal/webService/internal/controller/serviceController"
)

func Start() {
	r := gin.Default()
	r.POST("/login")
	r.GET("/home", serviceController.ListService)

	r.GET("/service", serviceController.ListService)
	r.POST("/service", serviceController.AddService)
	r.PUT("/service", serviceController.UpdateService)
	r.DELETE("/service", serviceController.DeleteService)

	r.Run(":8000")
}

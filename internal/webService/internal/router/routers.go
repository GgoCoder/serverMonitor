package router

import (
	"github.com/gin-gonic/gin"
	"serverMonitor/internal/webService/internal/controller/serviceController"
	"serverMonitor/internal/webService/internal/controller/loginController"
)

func Start() {
	r := gin.Default()
	r.POST("/login", LoginController.Login)
	r.POST("/regiter", LoginController.Regiter)
	

	r.GET("/home", serviceController.ListService)


	r.GET("/service", serviceController.ListService)
	r.POST("/service", serviceController.AddService)
	r.PUT("/service", serviceController.UpdateService)
	r.DELETE("/service", serviceController.DeleteService)

	r.Run(":8000")
}

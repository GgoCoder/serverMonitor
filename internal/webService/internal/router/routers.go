package router

import (
	"github.com/gin-gonic/gin"
	"serverMonitor/internal/webService/internal/controller/loginController"
	"serverMonitor/internal/webService/internal/controller/serviceController"
	"serverMonitor/internal/webService/internal/controller/testController"
)

func Start() {
	r := gin.Default()
	r.POST("/login", LoginController.Login)
	r.POST("/register", LoginController.Regiter)

	r.GET("/home", serviceController.ListService)

	r.GET("/service", serviceController.ListService)
	r.POST("/service", serviceController.AddService)
	r.PUT("/service", serviceController.UpdateService)
	r.DELETE("/service", serviceController.DeleteService)

	r.GET("/test", testController.TestGetApi)
	r.DELETE("/test", testController.TestDeleteApi)
	r.POST("/test", testController.TestPostApi)
	r.PUT("/test", testController.TestPutApi)

	r.Run(":8000")
}

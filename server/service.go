package server

import (
	"serverMonitor/asset"
	"serverMonitor/monitor"
	"time"

	"github.com/gin-gonic/gin"
)

func ListService(c *gin.Context) {
	services := monitor.ListService()
	ResponseWithJson(c, services)
}

func AddService(c *gin.Context) {
	service := &asset.Service{}
	service.CreateTime = time.Now()
	service.LastCheckTime = time.Now()
	err := c.BindJSON(service)
	if err != nil {
		ResponseWithError(c, "error body")
		return
	}
	monitor.AddService(service)
	ResponseWithJson(c, "post service successfully")
}

func UpdateService(c *gin.Context) {
	service := &asset.Service{}
	err := c.BindJSON(service)
	if err != nil {
		ResponseWithError(c, "error body")
		return
	}
	monitor.UpdateService(service)
	ResponseWithJson(c, "update service successfully")
}

func DeleteService(c *gin.Context) {
	service := &asset.Service{}
	err := c.BindJSON(service)
	if err != nil {
		ResponseWithError(c, "error body")
		return
	}
	monitor.DeleteService(service)
	ResponseWithJson(c, "delete service successfully")
}

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
	err = monitor.AddService(service)
	if err != nil{
		ResponseWithError(c, err)
		return
	}
	ResponseWithJson(c, "post service successfully")
}

func UpdateService(c *gin.Context) {
	service := &asset.Service{}
	err := c.BindJSON(service)
	if err != nil {
		ResponseWithError(c, "error body")
		return
	}
	err = monitor.UpdateService(service)	
	if err != nil{
		ResponseWithError(c, err)
		return
	}

	ResponseWithJson(c, "update service successfully")
}

func DeleteService(c *gin.Context) {
	service := &asset.Service{}
	err := c.BindJSON(service)
	if err != nil {
		ResponseWithError(c, "error body")
		return
	}
	err = monitor.DeleteService(service)
	if err != nil{
		ResponseWithError(c, err)
		return
	}
	ResponseWithJson(c, "delete service successfully")
}

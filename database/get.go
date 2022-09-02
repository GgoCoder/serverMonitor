package database

import (
	"serverMonitor/asset"
	"serverMonitor/global"
)

func ListService() []asset.Service {
	var services []asset.Service
	global.DB.Find(&services)
	return services
}

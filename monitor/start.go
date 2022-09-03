package monitor

import (
	"serverMonitor/database"
	"serverMonitor/global"
)

func Start() {
	global.MysqlUrl = "192.168.0.100"
	global.DB = database.Init()
	InitServiceRoot()
	go NewTimer()
}

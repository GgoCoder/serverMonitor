package start

import (
	"serverMonitor/internal/webService/internal/database"
	"serverMonitor/internal/webService/internal/monitor"
	"serverMonitor/internal/webService/internal/router"
)

func Start() {
	database.DbInit()
	router.Start()
	monitor.Start()
}

package main

import (
	"serverMonitor/monitor"
	"serverMonitor/server"
)

func main() {
	monitor.Start()
	server.Start()
	select {}
}

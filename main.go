package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"serverMonitor/monitor"
	"serverMonitor/server"
)

func main() {
	go func() {
		fmt.Printf("start to http pprof\n")
		http.ListenAndServe("0.0.0.0:8081", nil)
	}()
	monitor.Start()
	server.Start()
	select {}
}

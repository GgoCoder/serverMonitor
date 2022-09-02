package monitor

func Start() {
	InitServiceRoot()
	go NewTimer()
}

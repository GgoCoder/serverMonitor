package serviceregister

import (
	"serverMonitor/internal/serviceRegister/pkg/config"
	"serverMonitor/pkg/typed"
	"testing"
	"time"
)

func Test_Register(t *testing.T) {
	config.ServiceRegisterConfig = &typed.ConfigYaml{}
	config.ServiceRegisterConfig.Etcd.EndPoints = []string{"192.168.0.100:2379"}
	logGrpcService := &typed.MicroService{ServiceName: "logGrpcService", Addr: "127.0.0.1", Port: 20008}
	Register(logGrpcService.ServiceName, logGrpcService.Addr, logGrpcService.Port)
	go DiscoverServices(logGrpcService.ServiceName, nil)

	time.Sleep(5 * time.Second)
	logGrpcService.Addr = "192.168.0.100"
	logGrpcService.Port = 10000
	Register(logGrpcService.ServiceName, logGrpcService.Addr, logGrpcService.Port)
	time.Sleep(5 * time.Second)

	DeRegister(logGrpcService.ServiceName)
	time.Sleep(2 * time.Second)
}

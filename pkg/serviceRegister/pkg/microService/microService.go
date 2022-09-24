package microservice
import(
	"serverMonitor/pkg/typed"
	"serverMonitor/pkg/serviceRegister/pkg/register"
)
type HandleMicroService interface{
	Register(serviceName, addr string, port uint16, conf *typed.ConfigYaml)
	DeRegister(serviceName string, conf *typed.ConfigYaml) 
	DiscoverServices(serviceName string, service chan *typed.MicroService, conf *typed.ConfigYaml)
}

type MicroService typed.MicroService

func(m *MicroService)Register(){
	register.Register(m.ServiceName, m.Addr, m.Port, nil)
}

func(m *MicroService)DeRegister(){
	register.DeRegister(m.ServiceName, nil)
}

func(m *MicroService)DiscoverServices(){
	register.DiscoverServices(m.ServiceName, nil, nil)
}

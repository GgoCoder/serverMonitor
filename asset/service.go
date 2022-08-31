package asset


import (
	"os/exec"
	"sync"
	"fmt"
	"time"
)

type serviceStatus string

const(
	ServiceReady serviceStatus = "serviceReady"
	ServiceNoteReady serviceStatus = "serviceReady"
	
)

type Service struct {
	Status        serviceStatus
	Name          string
	CreateTime    time.Time
	LastCheckTime time.Time
	Url           string
}

type ServiceRoot struct{
	Services map[string]*Service
	sync.Mutex
}

var serviceRoot *ServiceRoot 

func checkServiceStatus(){
	serviceRoot.Lock()
	defer serviceRoot.Unlock()
	for name, service := range serviceRoot.Services{
		if isServiceReady(service){
			serviceRoot.Services[name].Status = ServiceReady
			continue
		}
		serviceRoot.Services[name].Status = ServiceNoteReady
	}
}

func NewTimer(){
	timer := time.NewTicker(time.Minute * 5)
	for {
		select{
		case <-timer.C:
			checkServiceStatus()
		}
	}
}

func isServiceReady(service * Service) bool {
	cmd :=  exec.Command("curl", service.Url)
	err := cmd.Run()
	if err != nil{
		fmt.Printf("service: %s is not reachable, err: %+v\n", service.Name, err)
		return false
	}
	return true
	
}

func addService(service *Service) {
	serviceRoot.Lock()
	defer serviceRoot.Unlock()
	serviceRoot.Services[service.Name] = service
}

func deleteService(service *Service) {
	serviceRoot.Lock()
	defer serviceRoot.Unlock()
	delete(serviceRoot.Services, service.Name)
}

func updateService(oldService, newService *Service) {
	serviceRoot.Lock()
	defer serviceRoot.Unlock()
	delete(serviceRoot.Services, oldService.Name)
	serviceRoot.Services[newService.Name] = newService
}

func initServiceRoot(){
	serviceRoot = &ServiceRoot{
		Services: make(map[string]*Service),
	}
}

func getServiceRoot()*ServiceRoot{
	if serviceRoot != nil{
		return serviceRoot
	}
	return nil
}

func AddService(service *Service){
	addService(service)
}

func UpdateService(oldService, newService *Service){
	updateService(oldService, newService)
}

func DeleteService(service *Service){
	deleteService(service)
}

func InitService(){
	initServiceRoot()
}

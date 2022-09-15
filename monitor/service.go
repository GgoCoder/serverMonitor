package monitor

import (
	"fmt"
	"os/exec"
	"serverMonitor/asset"
	"serverMonitor/database"
	"sync"
	"time"
)

type ServiceRoot struct {
	Services map[string]*asset.Service
	sync.Mutex
}

var serviceRoot *ServiceRoot

func checkServiceStatus() {
	serviceRoot.Lock()
	defer serviceRoot.Unlock()
	for name, service := range serviceRoot.Services {
		if isServiceReady(service) {
			serviceRoot.Services[name].Status = asset.ServiceReady
			continue
		}
		serviceRoot.Services[name].Status = asset.ServiceNotReady
	}
}

func NewTimer() {
	timer := time.NewTicker(time.Minute * 5)
	for {
		select {
		case <-timer.C:
			checkServiceStatus()
		}
	}
}

func isServiceReady(service *asset.Service) bool {
	cmd := exec.Command("curl", service.Url)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("service: %s is not reachable, err: %+v\n", service.Name, err)
		return false
	}
	return true

}

func AddService(service *asset.Service) error {
	serviceRoot.Lock()
	defer serviceRoot.Unlock()
	if _, ok := serviceRoot.Services[service.Name]; ok{
		return fmt.Errorf("service :%+v is existed\n", service.Name)
	}
	serviceRoot.Services[service.Name] = service
	err := database.InsertService(service)
	return err
}

func DeleteService(service *asset.Service) error {
	serviceRoot.Lock()
	defer serviceRoot.Unlock()
	delete(serviceRoot.Services, service.Name)
	err := database.DeleteService(service)
	return err
}

func UpdateService(newService *asset.Service) error {
	serviceRoot.Lock()
	defer serviceRoot.Unlock()
	var oldService *asset.Service
	if _, ok := serviceRoot.Services[newService.Name]; !ok {
		return fmt.Errorf("new service is not exist, update failed")
	}
	oldService = serviceRoot.Services[newService.Name]
	delete(serviceRoot.Services, oldService.Name)
	serviceRoot.Services[newService.Name] = newService
	err := database.UpdateService(oldService, newService)
	return err
}

func ListService() []asset.Service {
	serviceRoot.Lock()
	defer serviceRoot.Unlock()
	var services []asset.Service
	for _, s := range serviceRoot.Services {
		services = append(services, *s)
	}
	return services
}

func InitServiceRoot() {
	serviceRoot = &ServiceRoot{
		Services: make(map[string]*asset.Service),
	}
	services := database.ListService()
	for _, s := range services {
		serviceRoot.Services[s.Name] = &s
	}
	fmt.Printf("service root: %+v\n", serviceRoot)
}

func getServiceRoot() *ServiceRoot {
	if serviceRoot != nil {
		return serviceRoot
	}
	return nil
}

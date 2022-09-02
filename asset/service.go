package asset

import (
	"time"
)

type serviceStatus string

const (
	ServiceReady    serviceStatus = "serviceReady"
	ServiceNotReady serviceStatus = "serviceNotReady"
)

type Service struct {
	Status        serviceStatus
	Name          string
	CreateTime    time.Time
	LastCheckTime time.Time
	Url           string
}

func (o *Service) Copy(newService *Service) {
	o.Name = newService.Name
	o.Url = newService.Url
}

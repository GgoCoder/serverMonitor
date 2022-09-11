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
	Name          string
	CreateTime    time.Time
	LastCheckTime time.Time
	Url           string
	Status        serviceStatus
}

func (o *Service) Copy(newService *Service) {
	o.Name = newService.Name
	o.Url = newService.Url
}

func (Service) TableName() string {
	return "service"
}

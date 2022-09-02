package database

import (
	"fmt"
	"serverMonitor/asset"
	"serverMonitor/global"
)

func InsertService(service *asset.Service) error {
	result := global.DB.Create(service)
	if result.Error != nil {
		return fmt.Errorf("failed to insert service, err:%+v\n", result.Error)
	}
	return nil
}

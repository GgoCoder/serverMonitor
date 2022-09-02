package database

import (
	"fmt"
	"serverMonitor/asset"
	"serverMonitor/global"
)

func DeleteService(service *asset.Service) error {
	result := global.DB.Delete(service)
	if result.Error != nil {
		fmt.Errorf("failed to delete service, err:%+v\n", result.Error)
	}
	if result.RowsAffected > 1 {
		fmt.Errorf("delete service{%} more than 1", service.Name)
	}
	if result.RowsAffected == 0 {
		fmt.Errorf("service {%s} is not exist", service.Name)
	}
	return nil
}

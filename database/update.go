package database

import (
	"fmt"
	"serverMonitor/asset"
	"serverMonitor/global"
)

func UpdateService(oldService, newService *asset.Service) error {
	global.DB.First(oldService)
	oldService.Copy(newService)
	result := global.DB.Save(oldService)
	if result.Error != nil {
		fmt.Errorf("failed to update service{%s}, err:%+v\n", newService.Name, result.Error)
	}
	return nil
}

package database

import (
	"fmt"
	"serverMonitor/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//TODO(sync once)
func Init() *gorm.DB {
	dsn := fmt.Sprintf("grafana:grafana@tcp(%s:3306)/web_service?charset=utf8mb4&parseTime=True&loc=Local", global.MysqlUrl)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("open mysql err, err:  %+v", err))
	}
	return db
}

func PingDb() error {
	db := Init()
	client, err := db.DB()
	if err != nil {
		return err
	}
	if client.Ping() != nil {
		return fmt.Errorf("can not to connect to mysql")
	}
	return nil
}

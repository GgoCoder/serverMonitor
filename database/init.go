package database

import (
	"fmt"
	"serverMonitor/asset"
	"serverMonitor/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//TODO(sync once)
func Init(user, passwd, url, database string, port uint16) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, passwd, url, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("open mysql err, err:  %+v", err))
	}
	return db
}

func PingDb() error {
	db := global.DB
	client, err := db.DB()
	if err != nil {
		return err
	}
	if client.Ping() != nil {
		return fmt.Errorf("can not to connect to mysql")
	}
	return nil
}

func MigratorTable(tableName string) {
	switch tableName {
	case "service":
		global.DB.AutoMigrate(&asset.Service{})
	}
}

func init(){
	global.DB = Init(global.Config.Db.User, global.Config.Db.Passwd, global.Config.Db.URL, global.Config.Db.DbName, global.Config.Db.Port)
}

package global

import (
	"fmt"
	"serverMonitor/config"
	"serverMonitor/log"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var MysqlUrl string = "db"

var DB *gorm.DB

var Config *config.ConfigYaml

var Logger *zap.SugaredLogger
var Level = zap.NewAtomicLevel()
var ConfigPath = "/root/workspace/serverMonitor/config.yaml"

func init(){
	ReadConfig()
	Logger = log.InitLogger(Config.Log.FileName, Level)
}
func ReadConfig()*config.ConfigYaml{

	viper.SetConfigFile(ConfigPath)
	err := viper.ReadInConfig()
	if err != nil{
		fmt.Printf("failed to read config, err: %+v\n", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event){
		fmt.Printf("config change, reload config\n")
		reloadConfig()
	})
	reloadConfig()
	fmt.Printf("conf: %+v\n", Config)
	return nil
}

func reloadConfig(){
	Config = &config.ConfigYaml{
		Db: config.DbConfig{
			DbName: viper.GetString("db.dbName"),
			Database: viper.GetString("db.database"),
			URL: viper.GetString("db.url"),
			Port: viper.GetUint16("db.port"),
			User: viper.GetString("db.user"),
			Passwd: viper.GetString("db.passwd"),
		},
		Log: config.LogConfig{
			FileName: viper.GetString("log.fileName"),
			Level: viper.GetString("log.level"),
		},
	}
}

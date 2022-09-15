package global

import (
	"testing"
	"time"
)

func TestReadConfig(t *testing.T){
	ConfigPath = "./../config.yaml"
	ReadConfig()
	for {
		// fmt.Printf("%+v\n", Config)
		Logger.Infof("%+v\n", Config)
		Logger.Errorf("%+v\n", Config)
		time.Sleep(10 * time.Second)
	}
}

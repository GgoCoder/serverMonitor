package database

import (
	"fmt"
	"serverMonitor/asset"
	"serverMonitor/global"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	global.MysqlUrl = "192.168.0.100"
	assert.Equal(t, nil, PingDb())
}

func TestAdd(t *testing.T) {
	global.MysqlUrl = "192.168.0.100"
	global.DB = Init()
	MigratorTable("service")
	serviceMonitor := &asset.Service{}
	serviceMonitor.Name = "监控"
	serviceMonitor.Url = "monitor.rguo.top"
	serviceMonitor.CreateTime = time.Now()
	serviceMonitor.LastCheckTime = time.Now()
	serviceMonitor.Status = asset.ServiceNotReady
	fmt.Println(InsertService(serviceMonitor))
	serviceGit := &asset.Service{}
	serviceGit.Name = "Git"
	serviceGit.Url = "git.rguo.top"
	serviceGit.CreateTime = time.Now()
	serviceGit.LastCheckTime = time.Now()
	serviceGit.Status = asset.ServiceReady
	fmt.Println(InsertService(serviceGit))
}

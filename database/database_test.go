package database

import (
	"serverMonitor/asset"
	"serverMonitor/global"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	global.MysqlUrl = "192.168.0.100"
	assert.Equal(t, nil, PingDb())
}

func TestAdd(t *testing.T) {
	global.MysqlUrl = "192.168.0.100"
	global.DB = Init()
	serviceMonitor := &asset.Service{}
	serviceMonitor.Name = "监控"
	serviceMonitor.Url = "monitor.rguo.top"
	serviceMonitor.Status = asset.ServiceNotReady
	InsertService(serviceMonitor)
	serviceGit := &asset.Service{}
	serviceGit.Name = "Git"
	serviceGit.Url = "git.rguo.top"
	serviceGit.Status = asset.ServiceReady
	InsertService(serviceGit)

}

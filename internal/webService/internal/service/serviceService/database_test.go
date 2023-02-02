package serviceService

import (
	"fmt"
	"serverMonitor/pkg/typed"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	assert.Equal(t, nil, PingDb())
}

func TestAdd(t *testing.T) {
	MigratorTable("service")
	serviceMonitor := &typed.Service{}
	serviceMonitor.Name = "监控"
	serviceMonitor.Url = "xx.xx.xx"
	serviceMonitor.CreateTime = time.Now()
	serviceMonitor.LastCheckTime = time.Now()
	serviceMonitor.Status = typed.ServiceNotReady
	fmt.Println(InsertService(serviceMonitor))
	serviceGit := &typed.Service{}
	serviceGit.Name = "Git"
	serviceGit.Url = "yy.yy.yy"
	serviceGit.CreateTime = time.Now()
	serviceGit.LastCheckTime = time.Now()
	serviceGit.Status = typed.ServiceReady
	fmt.Println(InsertService(serviceGit))
}

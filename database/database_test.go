package database

import (
	"serverMonitor/global"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	global.MysqlUrl = "192.168.0.100"
	assert.Equal(t, nil, PingDb())
}

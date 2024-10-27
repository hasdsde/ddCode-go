package orm

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

const (
	host     = "124.70.56.154:3306"
	username = "tower"
	passwd   = "2024!@tower"
	db       = "tower"
)

func TestMysql(t *testing.T) {
	connConfig := &MysqlConfig{
		Host: host, Username: username, Passwd: passwd, DB: db,
		Timeout: time.Hour, MaxIdle: 10, MaxOpen: 100,
		IsDebug: true,
	}
	orm, err := NewOrm(connConfig, new(gorm.Config))
	if !assert.NoError(t, err) {
		return
	}
	if !assert.NoError(t, orm.Ping()) {
		return
	}
}

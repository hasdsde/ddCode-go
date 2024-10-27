package rbac

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"os"
)

type CasbinService struct {
	a *gormadapter.Adapter
	e *casbin.Enforcer
}

type CasbinServiceConfig struct {
	DriverName     string
	DataSourceName string
	TableName      string
	DbSpecified    []bool
	ModelFilePath  string
}

func InitRbacWithDB(db *gorm.DB, modelPath string) (*CasbinService, error) {
	_, err := os.Stat(modelPath)
	if err != nil {
		return nil, fmt.Errorf("文件不存在:%s", modelPath)
	}
	// 新版适配器表名为casbin_rule
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		return nil, fmt.Errorf("casbin加载失败:%s", err)
	}
	enforcer, err := casbin.NewEnforcer(modelPath, adapter)
	if err != nil {
		return nil, fmt.Errorf("casbin加载失败:%s", err)
	}
	err = enforcer.LoadPolicy()
	if err != nil {
		return nil, fmt.Errorf("casbin加载配置失败:%s", err)
	}
	return &CasbinService{adapter, enforcer}, nil
}

func InitTest() (*CasbinService, error) {
	a, err := gormadapter.NewAdapter("mysql", "tower:tower2024tower2024@tcp(124.70.56.154:3307)/tower", true)
	e, err := casbin.NewEnforcer("../../config/model.conf", a)
	return &CasbinService{a, e}, err
}

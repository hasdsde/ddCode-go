package global

import (
	"context"
	cache "ddCode-server/pkg/cache/redis"
	jwttoken "ddCode-server/pkg/jwt"
	"ddCode-server/pkg/rbac"
	orm "ddCode-server/pkg/storage/mysql"
	"ddCode-server/pkg/storage/oss"
	logger "ddCode-server/pkg/zlogs"
	"flag"
	"fmt"
	"gorm.io/gorm"
	"runtime"
	"strings"
	"time"
)

func GetRootPath() string {
	_, filename, _, _ := runtime.Caller(0)
	return strings.Replace(filename, "/global/test.go", "", 1)
}
func SetUpTest() error {
	// 项目目录
	rootPath := GetRootPath()
	configFile := flag.String("f", rootPath+"/config/config.yaml", "the config file")
	configModeFile := flag.String("m", "dev", "the mode")
	// 解析yaml
	flag.Parse()
	//filename, _ := filepath.Abs(*configFile)
	c, err := Parse(configFile, configModeFile)
	if err != nil {
		fmt.Printf("解析配置文件失败:%v\n", err)
		return err
	}

	// 初始化日志
	logging, err := logger.NewLogger(c.Mode, c.Log)
	if err != nil {
		fmt.Printf("初始化日志失败:%v\n", err)
		return err
	}

	// 初始化GORM Client
	ormCli, err := orm.NewOrm(
		&orm.MysqlConfig{
			Host: c.Mysql.Host, Username: c.Mysql.User, Passwd: c.Mysql.Pass, Port: c.Mysql.Port, DB: c.Mysql.Database,
			Timeout: time.Hour, MaxIdle: 10, MaxOpen: 10,
			IsDebug: true,
		},
		new(gorm.Config))
	if err != nil {
		logging.Errorf("建立数据连接失败:%v\n", err)
		return err
	}

	// 加载casbin
	rbacService, err := rbac.InitRbacWithDB(ormCli.DB(), rootPath+c.Casbin.FilePath)
	if err != nil {
		logging.Errorf("初始化角色权限失败:%v\n", err)
		return err
	}

	// Jwt
	tokenMaker, err := jwttoken.NewJWTMaker(c.Jwt.Signature, c.Jwt.Duration)
	if err != nil {
		logging.Fatalf("初始化token失败: %v", err)
		return err
	}
	// 初始化redis Client
	redisCli, err := cache.NewClient(context.Background(), c.Redis.Addr, c.Redis.DB, cache.WithPasswd(c.Redis.Pass))
	if err != nil {
		logging.Errorf("init redis Error", err)
		return err
	}

	// 初始化 oss
	ossCli, err := oss.NewOSS(c.Oss.Url, c.Oss.AccessKey, c.Oss.SecretKey, c.Oss.Bucket)
	if err != nil {
		logging.Errorf("init minio Error", err)
		return err
	}

	Logger = logging
	Db = ormCli
	JwtMaker = tokenMaker
	Rbac = rbacService
	Redis = redisCli
	Oss = ossCli

	return nil
}

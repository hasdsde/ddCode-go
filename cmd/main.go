package main

import (
	"context"
	_ "ddCode-server/docs/swagger"
	"ddCode-server/global"
	cache "ddCode-server/pkg/cache/redis"
	jwttoken "ddCode-server/pkg/jwt"
	"ddCode-server/pkg/rbac"
	orm "ddCode-server/pkg/storage/mysql"
	"ddCode-server/pkg/storage/oss"
	logger "ddCode-server/pkg/zlogs"
	"ddCode-server/router"
	"flag"
	"fmt"
	"gorm.io/gorm"
	"time"
)

var configFile = flag.String("f", "./config/config.yaml", "the config file")
var configModeFile = flag.String("m", "dev", "the mode")

// @title ddCode-server
// @version 1.0
// @description ddCode-server
// @host localhost:8070
// @basePath /api/v1
func main() {
	// 解析yaml
	flag.Parse()
	//filename, _ := filepath.Abs(*configFile)
	c, err := global.Parse(configFile, configModeFile)
	if err != nil {
		fmt.Printf("解析配置文件失败:%v\n", err)
		return
	}

	// 初始化日志
	logging, err := logger.NewLogger(c.Mode, c.Log)
	if err != nil {
		fmt.Printf("初始化日志失败:%v\n", err)
		return
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
		return
	}

	// 加载casbin
	rbacService, err := rbac.InitRbacWithDB(ormCli.DB(), c.Casbin.FilePath)
	if err != nil {
		logging.Errorf("初始化角色权限失败:%v\n", err)
		return
	}

	// Jwt
	tokenMaker, err := jwttoken.NewJWTMaker(c.Jwt.Signature, c.Jwt.Duration)
	if err != nil {
		logging.Fatalf("初始化token失败: %v", err)
		return
	}
	// 初始化redis Client
	redisCli, err := cache.NewClient(context.Background(), c.Redis.Addr, c.Redis.DB, cache.WithPasswd(c.Redis.Pass))
	if err != nil {
		logging.Errorf("init redis Error", err)
		return
	}

	// 初始化 oss
	ossCli, err := oss.NewOSS(c.Oss.Url, c.Oss.AccessKey, c.Oss.SecretKey, c.Oss.Bucket)
	if err != nil {
		logging.Errorf("init minio Error", err)
		return
	}
	//// 初始化 es
	//esCli, err := es.NewElastic(&es.Config{Addrs: c.ESConf.Addrs, User: c.ESConf.User, Pwd: c.ESConf.Pwd}, es.WithLogger(logging))
	//if err != nil {
	//	logging.Errorf("init Es Error", err)
	//	return
	//}
	//// 初始化 bloom filter
	//bloom, err := bloomfilter.NewBloom(context.Background(), 10000, 0.0001)
	//if err != nil {
	//	logging.Errorf("init bloomfilter Error:%v", err)
	//	return
	//}
	//// 初始化 levelDB
	//levelDB, err := leveldb.NewLevelDB(&leveldb.Config{
	//	FilePath: c.LevelDB.Path,
	//})
	//if err != nil {
	//	logging.Fatalf("init level db error: %v", err)
	//}
	//gspCli, err := oss.NewGSP(c.GSP.Addr, c.GSP.AccessKey, c.GSP.SecretKey, c.GSP.Regions)
	//if err != nil {
	//	logging.Errorf("init oss Error", err)
	//	return
	//}

	global.Logger = logging
	global.Db = ormCli
	global.JwtMaker = tokenMaker
	global.Rbac = rbacService
	global.Redis = redisCli
	global.Oss = ossCli

	// 初始化服务端
	r := router.Init(c.Mode)
	err = r.Run(fmt.Sprintf("%s:%d", c.Host, c.Port))
	if err != nil {
		logging.Fatalf("初始化服务端失败: %v", err)
		return
	}
}

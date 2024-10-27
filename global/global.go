package global

import (
	cache "ddCode-server/pkg/cache/redis"
	jwttoken "ddCode-server/pkg/jwt"
	"ddCode-server/pkg/rbac"
	orm "ddCode-server/pkg/storage/mysql"
	"ddCode-server/pkg/storage/oss"
	logger "ddCode-server/pkg/zlogs"
)

var (
	Db       *orm.Client
	Logger   *logger.AppLogger
	JwtMaker jwttoken.Maker
	Rbac     *rbac.CasbinService
	Redis    cache.Redis
	Oss      *oss.OSS
)

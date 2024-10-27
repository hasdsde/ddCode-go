package cache

import (
	"context"
	"time"
)

// 声明对list的操作点
type ListFlag int8

const (
	StartPoint ListFlag = iota + 1 // 列表的头部（左边）
	EndPoint                       // 列表的尾部（右边）
)

// 声明支持的redis部署模式
type DeployMode string

const (
	ClientMod   DeployMode = "client"   // 单机模式
	ClusterMod  DeployMode = "cluster"  // 集群模式
	SentinelMod DeployMode = "sentinel" // 哨兵模式
)

type Message struct {
	Channel      string
	Pattern      string
	Payload      interface{}
	PayloadSlice []string
}

// ZSetMember 有序集合的数据结构
type ZSetMember struct { // nolint
	Score  float64     `json:"score"`
	Member interface{} `json:"member"`
}
type Redis interface {
	IsExist(ctx context.Context, key ...string) bool
	Del(ctx context.Context, key ...string) error
	SetExpire(ctx context.Context, key string, ttl time.Duration) error
	GetMixed(ctx context.Context, key string, value interface{}) error

	SetStr(ctx context.Context, key, value string) error
	SetStrTTL(ctx context.Context, key, value string, ttl time.Duration) error
	SetNX(ctx context.Context, key, value string, ttl time.Duration) error

	SetHash(ctx context.Context, key string, value map[string]interface{}) error
	GetHashField(ctx context.Context, key, field string) (string, error)

	PushList(ctx context.Context, key string, values ...interface{}) error
	LenList(ctx context.Context, key string) (int64, error)
	PopList(ctx context.Context, key string, value interface{}) error

	AddSet(ctx context.Context, key string, values ...interface{}) error
	CheckSetMember(ctx context.Context, key string, value interface{}) (bool, error)
	RemSetEle(ctx context.Context, key string, values ...interface{}) error

	AddZSet(ctx context.Context, key string, members ...*ZSetMember) error
	CardZSet(ctx context.Context, key string) (int64, error)
	MembersWithScoreZSet(ctx context.Context, key string) ([]*ZSetMember, error)
	RemMembersZSet(ctx context.Context, key string, members ...string) error
	CountRedis(ctx context.Context, key string) (int64, error)

	Publish(ctx context.Context, channel string, message interface{}) error
	Subscribe(ctx context.Context, channels ...string) (<-chan *Message, error)
}

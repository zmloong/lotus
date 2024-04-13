package redis

import (
	"time"

	"github.com/zmloong/lotus/utils/mapstructure"
)

type RedisType int8

const (
	Redis_Single = iota
	Redis_Cluster
)

// /redis 存储数据格式化类型
type RedisStorageType int8

const (
	JsonData RedisStorageType = iota
	ProtoData
)

type Option struct {
	RedisType              RedisType
	Redis_Single_Addr      string
	Redis_Single_Password  string
	Redis_Single_DB        int
	Redis_Single_PoolSize  int
	Redis_Cluster_Addr     []string
	Redis_Cluster_Password string
	RedisStorageType       RedisStorageType
	TimeOut                time.Duration
}
type Optionfn func(*Option)

func SetRedisType(v RedisType) Optionfn {
	return func(o *Option) {
		o.RedisType = v
	}
}

// /RedisUrl = "127.0.0.1:6379"
func SetRedis_Single_Addr(v string) Optionfn {
	return func(o *Option) {
		o.Redis_Single_Addr = v
	}
}
func SetRedis_Single_Password(v string) Optionfn {
	return func(o *Option) {
		o.Redis_Single_Password = v
	}
}
func SetRedis_Single_DB(v int) Optionfn {
	return func(o *Option) {
		o.Redis_Single_DB = v
	}
}

func SetRedis_Single_PoolSize(v int) Optionfn {
	return func(o *Option) {
		o.Redis_Single_PoolSize = v
	}
}
func Redis_Cluster_Addr(v []string) Optionfn {
	return func(o *Option) {
		o.Redis_Cluster_Addr = v
	}
}

func SetRedis_Cluster_Password(v string) Optionfn {
	return func(o *Option) {
		o.Redis_Cluster_Password = v
	}
}
func SetRedisStorageType(v RedisStorageType) Optionfn {
	return func(o *Option) {
		o.RedisStorageType = v
	}
}

func SetTimeOut(v time.Duration) Optionfn {
	return func(o *Option) {
		o.TimeOut = v
	}
}

// 赋默认值
func newOptions(config map[string]interface{}, optfns ...Optionfn) Option {
	option := Option{
		Redis_Single_Addr:      "127.0.0.1:6379",
		Redis_Single_Password:  "",
		Redis_Single_DB:        1,
		Redis_Cluster_Addr:     []string{"127.0.0.1:6379"},
		Redis_Cluster_Password: "",
		TimeOut:                time.Second * 3,
		Redis_Single_PoolSize:  100,
	}
	if config != nil {
		mapstructure.Decode(config, &option)
	}
	for _, o := range optfns {
		o(&option)
	}
	return option
}

// 赋默认值
func newOptionsByOption(optfns ...Optionfn) Option {
	option := Option{
		Redis_Single_Addr:      "127.0.0.1:6379",
		Redis_Single_Password:  "",
		Redis_Single_DB:        1,
		Redis_Cluster_Addr:     []string{"127.0.0.1:6379"},
		Redis_Cluster_Password: "",
		TimeOut:                time.Second * 3,
		Redis_Single_PoolSize:  100,
	}
	for _, o := range optfns {
		o(&option)
	}
	return option
}

type RMutexOption struct {
	expiry int
	delay  time.Duration
}
type RMutexOptionfn func(*RMutexOption)

func SetExpiry(v int) RMutexOptionfn {
	return func(o *RMutexOption) {
		o.expiry = v
	}
}
func SetDelay(v time.Duration) RMutexOptionfn {
	return func(o *RMutexOption) {
		o.delay = v
	}
}

func newRMutexOptions(optfns ...RMutexOptionfn) RMutexOption {
	option := RMutexOption{
		expiry: 5,
		delay:  time.Millisecond * 50,
	}

	for _, o := range optfns {
		o(&option)
	}
	return option
}

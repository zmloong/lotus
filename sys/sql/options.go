package sql

import (
	"time"

	"github.com/zmloong/lotus/sys/log"
	"github.com/zmloong/lotus/utils/mapstructure"
)

type SqlType string

const (
	SqlServer SqlType = "sqlserver"
	MySql     SqlType = "mysql"
	Oracle    SqlType = "godror"
	DM        SqlType = "dm"
	PG        SqlType = "postgres"
)

type Option func(*Options)
type Options struct {
	SqlType SqlType
	SqlUrl  string
	TimeOut time.Duration
}

func SetSqlType(v SqlType) Option {
	return func(o *Options) {
		o.SqlType = v
	}
}

func SetSqlUrl(v string) Option {
	return func(o *Options) {
		o.SqlUrl = v
	}
}

func SetTimeOut(v time.Duration) Option {
	return func(o *Options) {
		o.TimeOut = v
	}
}

func newOptions(config map[string]interface{}, opts ...Option) Options {
	options := Options{
		SqlType: MySql,
		TimeOut: 3 * time.Second,
	}
	if config != nil {
		mapstructure.Decode(config, &options)
	}
	for _, o := range opts {
		o(&options)
	}
	if len(options.SqlUrl) == 0 {
		log.Errorf("start sqls Missing necessary configuration : SqlUrl is nul")
	}
	return options
}

func newOptionsByOption(opts ...Option) Options {
	options := Options{
		TimeOut: 3 * time.Second,
	}
	for _, o := range opts {
		o(&options)
	}
	if len(options.SqlUrl) == 0 {
		log.Errorf("start sqls Missing necessary configuration : SqlUrl is nul")
	}
	return options
}

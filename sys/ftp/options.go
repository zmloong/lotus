package ftp

import (
	"github.com/zmloong/lotus/utils/mapstructure"
)

type Option struct {
	ServerIp     string // 访问ip地址
	Port         int32  // 访问端口
	UserName     string // 登录用户名
	PassWord     string // 登录密码
	Directory    string // 访问目录
	TimeOut      int    // 超时时间 默认5s
	Interval     string // 采集间隔时间
	RegularRules string // 文件过滤正则规则
}
type Optionfn func(*Option)

func SetServerIp(v string) Optionfn {
	return func(o *Option) {
		o.ServerIp = v
	}
}
func SetPort(v int32) Optionfn {
	return func(o *Option) {
		o.Port = v
	}
}
func SetUsername(v string) Optionfn {
	return func(o *Option) {
		o.UserName = v
	}
}
func SetPassword(v string) Optionfn {
	return func(o *Option) {
		o.PassWord = v
	}
}
func SetDirectory(v string) Optionfn {
	return func(o *Option) {
		o.Directory = v
	}
}
func SetTimeout(v int) Optionfn {
	return func(o *Option) {
		o.TimeOut = v
	}
}
func SetInterval(v string) Optionfn {
	return func(o *Option) {
		o.Interval = v
	}
}
func SetRegularrules(v string) Optionfn {
	return func(o *Option) {
		o.RegularRules = v
	}
}

func newOptions(config map[string]interface{}, optfns ...Optionfn) Option {
	option := Option{
		Port:    21,
		TimeOut: 5,
	}
	if config != nil {
		mapstructure.Decode(config, &option)
	}
	for _, o := range optfns {
		o(&option)
	}
	return option
}
func newOptionsByOptionFn(optfns ...Optionfn) Option {
	option := Option{
		Port:    21,
		TimeOut: 5,
	}
	for _, o := range optfns {
		o(&option)
	}
	return option
}

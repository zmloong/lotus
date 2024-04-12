package ftp

import (
	"github.com/zmloong/lotus/utils/mapstructure"
)

type Option struct {
	serverip     string // 访问ip地址
	port         int32  // 访问端口
	username     string // 登录用户名
	password     string // 登录密码
	directory    string // 访问目录
	timeout      int    // 超时时间 默认5s
	interval     string // 采集间隔时间
	regularrules string // 文件过滤正则规则
}
type Optionfn func(*Option)

func SetServerIp(v string) Optionfn {
	return func(o *Option) {
		o.serverip = v
	}
}
func SetPort(v int32) Optionfn {
	return func(o *Option) {
		o.port = v
	}
}
func SetUsername(v string) Optionfn {
	return func(o *Option) {
		o.username = v
	}
}
func SetPassword(v string) Optionfn {
	return func(o *Option) {
		o.password = v
	}
}
func SetDirectory(v string) Optionfn {
	return func(o *Option) {
		o.directory = v
	}
}
func SetTimeout(v int) Optionfn {
	return func(o *Option) {
		o.timeout = v
	}
}
func SetInterval(v string) Optionfn {
	return func(o *Option) {
		o.interval = v
	}
}
func SetRegularrules(v string) Optionfn {
	return func(o *Option) {
		o.regularrules = v
	}
}

func newOptions(config map[string]interface{}, optfns ...Optionfn) Option {
	option := Option{
		port:    21,
		timeout: 5,
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
		port:    21,
		timeout: 5,
	}
	for _, o := range optfns {
		o(&option)
	}
	return option
}

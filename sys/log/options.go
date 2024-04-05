package log

import (
	"github.com/zmloong/lotus/utils/mapstructure"
)

type Loglevel int8

const (
	DebugLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

type LogEncoder int8

const (
	Console LogEncoder = iota
	JSON
)

type Options struct {
	FileName      string     //日志文件名包含
	Loglevel      Loglevel   //日志输出级别
	Debugmode     bool       //是否debug模式
	Encoder       LogEncoder //日志输出样式
	Loglayer      int        //日志堆栈信息打印层级
	LogMaxSize    int        //每个日志文件最大尺寸 单位 M 默认 1024M
	LogMaxBackups int        //最多保留备份个数	默认 10个
	LogMaxAge     int        //文件最多保存多少天 默认 7天
}

type Optionfn func(*Options)

func SetFileName(v string) Optionfn {
	return func(o *Options) {
		o.FileName = v
	}
}
func SetLoglevel(v Loglevel) Optionfn {
	return func(o *Options) {
		o.Loglevel = v
	}
}

func SetDebugMode(v bool) Optionfn {
	return func(o *Options) {
		o.Debugmode = v
	}
}
func SetEncoder(v LogEncoder) Optionfn {
	return func(o *Options) {
		o.Encoder = v
	}
}
func SetLoglayer(v int) Optionfn {
	return func(o *Options) {
		o.Loglayer = v
	}
}
func SetLogMaxSize(v int) Optionfn {
	return func(o *Options) {
		o.LogMaxSize = v
	}
}
func SetLogMaxBackups(v int) Optionfn {
	return func(o *Options) {
		o.LogMaxBackups = v
	}
}
func SetLogMaxAge(v int) Optionfn {
	return func(o *Options) {
		o.LogMaxAge = v
	}
}

func newOptions(config map[string]interface{}, opts ...Optionfn) Options {
	options := Options{
		Loglevel:      WarnLevel,
		Debugmode:     false,
		Encoder:       Console,
		Loglayer:      2,
		LogMaxSize:    1024,
		LogMaxBackups: 10,
		LogMaxAge:     7,
	}
	if config != nil {
		mapstructure.Decode(config, &options)
	}
	for _, o := range opts {
		o(&options)
	}
	return options
}
func newOptionsByOptionFn(opts ...Optionfn) Options {
	options := Options{
		Loglevel:      WarnLevel,
		Debugmode:     false,
		Encoder:       Console,
		Loglayer:      2,
		LogMaxSize:    1024,
		LogMaxBackups: 10,
		LogMaxAge:     7,
	}

	for _, o := range opts {
		o(&options)
	}
	return options
}

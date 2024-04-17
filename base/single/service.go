package single

import (
	"fmt"

	"github.com/zmloong/lotus/core"
	"github.com/zmloong/lotus/core/cbase"
	"github.com/zmloong/lotus/sys/event"
	"github.com/zmloong/lotus/sys/log"
)

type SingleService struct {
	cbase.ServiceBase
	opts *Options
}

func (this *SingleService) GetId() string {
	return this.opts.Setting.Id
}

func (this *SingleService) GetType() string {
	return this.opts.Setting.Type
}

func (this *SingleService) GetVersion() string {
	return this.opts.Version
}
func (this *SingleService) GetIp() string {
	return this.opts.Setting.Ip
}
func (this *SingleService) GetSettings() core.ServiceSttings {
	return this.opts.Setting
}

func (this *SingleService) Configure(opts ...Option) {
	this.opts = newOptions(opts...)
}

// 初始化服务系统模块
func (this *SingleService) InitSys() {
	if err := log.OnInit(this.opts.Setting.Sys["log"]); err != nil {
		panic(fmt.Sprintf("初始化log系统失败 %s", err.Error()))
	}
	if err := event.OnInit(this.opts.Setting.Sys["event"]); err != nil {
		log.Panicf(fmt.Sprintf("初始化event系统失败 %s", err.Error()))
	}
}

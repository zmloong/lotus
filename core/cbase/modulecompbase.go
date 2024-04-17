package cbase

import (
	"github.com/zmloong/lotus/core"
)

type ModuleCompBase struct{}

func (this *ModuleCompBase) Init(service core.IService, module core.IModule, comp core.IModuleComp, options core.IModuleOptions) (err error) {
	return
}

func (this *ModuleCompBase) Start() (err error) {
	return
}

func (this *ModuleCompBase) Destroy() (err error) {
	return
}

package cbase

import (
	"github.com/zmloong/lotus/core"
	"github.com/zmloong/lotus/utils/mapstructure"
)

type ServiceCompOptions struct {
	core.ICompOptions
}

func (this *ServiceCompOptions) LoadConfig(settings map[string]interface{}) (err error) {
	if settings != nil {
		mapstructure.Decode(settings, &this)
	}
	return
}

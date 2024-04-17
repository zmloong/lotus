package cbase

import "github.com/zmloong/lotus/utils/mapstructure"

type ModuleOptions struct {
}

func (this *ModuleOptions) LoadConfig(settings map[string]interface{}) (err error) {
	if settings != nil {
		mapstructure.Decode(settings, &this)
	}
	return
}

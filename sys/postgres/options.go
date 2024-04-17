package postgres

import (
	"github.com/zmloong/lotus/utils/mapstructure"
)

type Option func(*Options)

type Options struct{}

func newOptions(config map[string]interface{}, opt ...Option) Options {
	options := Options{}
	if config != nil {
		mapstructure.Decode(config, &options)
	}
	for _, o := range opt {
		o(&options)
	}
	return options
}
func newOptionsByOption(opt ...Option) Options {
	options := Options{}
	for _, o := range opt {
		o(&options)
	}
	return options
}

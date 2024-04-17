package rpc

import (
	"github.com/zmloong/lotus/utils/mapstructure"
)

type Option func(*Options)
type Options struct {
	RPCConnType   RPCConnType
	Listener      RPCListener
	ClusterTag    string
	ServiceId     string
	MaxCoroutine  int
	RpcExpired    int
	Nats_Addr     string
	Kafka_Host    []string
	Kafka_Version string
}

func SetClusterTag(v string) Option {
	return func(o *Options) {
		o.ClusterTag = v
	}
}
func SetServiceId(v string) Option {
	return func(o *Options) {
		o.ServiceId = v
	}
}
func newOptions(config map[string]interface{}, opts ...Option) Options {
	options := Options{
		RPCConnType:   Nats,
		MaxCoroutine:  2000,
		RpcExpired:    5,
		Kafka_Version: "1.0.0",
	}
	if config != nil {
		mapstructure.Decode(config, &options)
	}
	for _, o := range opts {
		o(&options)
	}
	return options
}

func newOptionsByOption(opts ...Option) Options {
	options := Options{
		RPCConnType:   Nats,
		MaxCoroutine:  2000,
		RpcExpired:    5,
		Kafka_Version: "1.0.0",
	}
	for _, o := range opts {
		o(&options)
	}
	return options
}

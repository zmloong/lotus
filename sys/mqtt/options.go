package mqtt

import (
	"time"

	"github.com/zmloong/lotus/utils/container/id"
	_ "github.com/zmloong/lotus/utils/container/id"
	"github.com/zmloong/lotus/utils/mapstructure"
)

type Options struct {
	MqttAddr string
	MqttPort int
	ClientID string
	UserName string
	PassWord string
	MqttDoor IMqttDoor
}
type Optionfn func(*Options)

func SetMqttAddr(v string) Optionfn {
	return func(o *Options) {
		o.MqttAddr = v
	}
}
func SetMqttPort(v int) Optionfn {
	return func(o *Options) {
		o.MqttPort = v
	}
}

func SetClientID(v string) Optionfn {
	return func(o *Options) {
		o.ClientID = v
	}
}

func SetUserName(v string) Optionfn {
	return func(o *Options) {
		o.UserName = v
	}
}
func SetMPassword(v string) Optionfn {
	return func(o *Options) {
		o.PassWord = v
	}
}

func SetMqttDoor(v IMqttDoor) Optionfn {
	return func(o *Options) {
		o.MqttDoor = v
	}
}
func newOptions(config map[string]interface{}, optfns ...Optionfn) Options {
	options := Options{
		ClientID: "lotus_mqtt_id" + time.Now().Format("2006-01-02-15:04:05"),
		MqttDoor: new(DefMqttDoor),
	}
	if config != nil {
		mapstructure.Decode(config, &options)
	}
	for _, o := range optfns {
		o(&options)
	}

	return options
}
func newOptionsByOptionFn(optfns ...Optionfn) Options {
	options := Options{
		ClientID: "lotus_mqtt_" + id.LotusUid(),
		MqttDoor: new(DefMqttDoor),
	}

	for _, o := range optfns {
		o(&options)
	}

	return options
}

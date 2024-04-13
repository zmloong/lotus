package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type (
	IMqtt interface {
		Subscribe(topic string, qos byte, callback mqtt.MessageHandler) mqtt.Token
		Publish(topic string, qos byte, retained bool, payload interface{}) mqtt.Token
	}
	IMqttDoor interface {
		ConnectHandler(client mqtt.Client)
		ConnectLostHandler(client mqtt.Client, err error)
		MessageHandler(client mqtt.Client, msg mqtt.Message)
	}
)

var (
	defsys IMqtt
)

func OnInit(config map[string]interface{}, option ...Optionfn) (err error) {
	defsys, err = newSys(newOptions(config, option...))
	return
}
func NewSys(option ...Optionfn) (sys IMqtt, err error) {
	sys, err = newSys(newOptionsByOptionFn(option...))
	return
}

func Subscribe(topic string, qos byte, callback mqtt.MessageHandler) mqtt.Token {
	return defsys.Subscribe(topic, qos, callback)
}

func Publish(topic string, qos byte, retained bool, payload interface{}) mqtt.Token {
	return defsys.Publish(topic, qos, retained, payload)
}

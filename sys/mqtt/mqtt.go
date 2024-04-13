package mqtt

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func newSys(options Options) (sys *Mqtt, err error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", options.MqttAddr, options.MqttPort))
	opts.SetClientID(options.ClientID)
	opts.SetUsername(options.UserName)
	opts.SetPassword(options.PassWord)
	opts.SetDefaultPublishHandler(options.MqttDoor.MessageHandler)
	opts.OnConnect = options.MqttDoor.ConnectHandler
	opts.OnConnectionLost = options.MqttDoor.ConnectLostHandler
	client := mqtt.NewClient(opts)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		err = token.Error()
		return
	}
	sys = &Mqtt{
		client: client,
		token:  token,
	}
	return
}

type Mqtt struct {
	client mqtt.Client
	token  mqtt.Token
}

func (z *Mqtt) Subscribe(topic string, qos byte, callback mqtt.MessageHandler) mqtt.Token {
	return z.client.Subscribe(topic, qos, callback)
}

func (z *Mqtt) Publish(topic string, qos byte, retained bool, payload interface{}) mqtt.Token {
	return z.client.Publish(topic, qos, retained, payload)
}

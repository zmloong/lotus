package mqtt

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type DefMqttDoor struct{}

// 连接的回掉函数
func (z *DefMqttDoor) ConnectHandler(client mqtt.Client) {
	fmt.Printf("ConnectHandler: succ \n")
}

// 丢失连接的回掉函数
func (z *DefMqttDoor) ConnectLostHandler(client mqtt.Client, err error) {
	fmt.Printf("ConnectLostHandler: succ \n")
}

// 创建全局  mqtt 消息处理 handler
func (z *DefMqttDoor) MessageHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("MessageHandler: topic:%s message:%v \n", msg.Topic(), msg.Payload())
}

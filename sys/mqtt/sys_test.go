package mqtt_test

import (
	"fmt"
	"testing"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	lsmqtt "github.com/zmloong/lotus/sys/mqtt"
)

func Test_sys(t *testing.T) {
	// sub(client)
	if err := lsmqtt.OnInit(map[string]interface{}{
		"MqttAddr": "118.178.94.244",
		"MqttPort": 1883,
		"UserName": "",
		"PassWord": "",
	}); err != nil {
		fmt.Printf("start sys err:%v", err)
	} else {
		sub()
		publish()
	}

}

func publish() {
	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("Message %d", i)
		token := lsmqtt.Publish("testtopic/1", 0, false, text)
		token.Wait()
		time.Sleep(time.Second)
	}
}

func sub() {
	token := lsmqtt.Subscribe("testtopic/1", 1, func(c mqtt.Client, m mqtt.Message) {
		fmt.Printf("testtopic/1 %d \n", m.Payload())
	})
	token.Wait()
	fmt.Printf("Subscribed to topic: %s \n", "testtopic/1")
}

package event_test

import (
	"fmt"
	"testing"

	"github.com/zmloong/lotus/sys/event"
)

func Test_sys(t *testing.T) {
	if err := event.OnInit(nil); err == nil {
		event.Register(event.Event_Key("TestEvent"), func() {
			fmt.Printf("TestEvent TriggerEvent")
		})
		event.TriggerEvent(event.Event_Key("TestEvent"))
	}
}

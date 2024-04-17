package event

import (
	"reflect"

	"github.com/zmloong/lotus/sys/log"
)

type Event_Key string //事件Key
type (
	FunctionInfo struct {
		Function  reflect.Value
		Goroutine bool
	}
	IEventSys interface {
		Register(eId Event_Key, f interface{}) (err error)
		RegisterGO(eId Event_Key, f interface{}) (err error)
		RemoveEvent(eId Event_Key, f interface{}) (err error)
		TriggerEvent(eId Event_Key, agr ...interface{})
	}
)

var (
	defsys IEventSys
)

func OnInit(config map[string]interface{}, option ...Option) (err error) {
	defsys, err = newSys(newOptions(config, option...))
	return
}

func NewSys(option ...Option) (err error) {
	defsys, err = newSys(newOptionsByOption(option...))
	return
}

func Register(eId Event_Key, f interface{}) (err error) {
	return defsys.Register(eId, f)
}

func RegisterGO(eId Event_Key, f interface{}) (err error) {
	return defsys.Register(eId, f)
}

func RemoveEvent(eId Event_Key, f interface{}) (err error) {
	return defsys.RemoveEvent(eId, f)
}

func TriggerEvent(eId Event_Key, agr ...interface{}) {
	if defsys != nil {
		defsys.TriggerEvent(eId, agr...)
	} else {
		log.Warnf("event no start")
	}
}

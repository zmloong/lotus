package lotus

import (
	"log"
	"runtime"
)

// 错误采集
func Recover(tag string) {
	if r := recover(); r != nil {
		buf := make([]byte, 1024)
		l := runtime.Stack(buf, false)
		log.Panicf("%s - %v: %s", tag, r, buf[:l])
	}
}

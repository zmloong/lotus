package id

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/xid"
)

var lastTimestamp int64 = time.Now().Unix()
var count byte = 0

func LotusUid() string {
	return strings.ToUpper(fmt.Sprintf("%x", getUUIDBytes())[:13])
}
func getUUIDBytes() []byte {
	uuidBytes := make([]byte, 7)
	offset := 0
	copy(uuidBytes[offset:], getPidBytes())
	offset += 2
	copy(uuidBytes[offset:], getTimestampBytes())
	offset += 4
	copy(uuidBytes[offset:], getCountInTimeBytes())
	return uuidBytes
}

// 获取时间戳内偏移量
func getCountInTimeBytes() []byte {
	bytes := make([]byte, 1)
	now := time.Now().Unix()
	if now == lastTimestamp {
		count++
	} else {
		count = 0
	}
	lastTimestamp = now
	bytes[0] = count
	return bytes
}

// 获取时间戳的byte[]数组
func getTimestampBytes() []byte {
	bytes := make([]byte, 4)
	timestamp := uint32(time.Now().Unix())
	for i := 3; i >= 0; i-- {
		bytes[i] = byte(timestamp)
		timestamp >>= 8
	}
	return bytes
}

// 获取本程序的进程号byte[]数组
func getPidBytes() []byte {
	bytes := make([]byte, 2)
	pid := uint16(os.Getpid())
	bytes[0] = byte(pid >> 8)
	bytes[1] = byte(pid)
	return bytes
}

// xid
func NewXId() string {
	id := xid.New()
	return id.String()
}

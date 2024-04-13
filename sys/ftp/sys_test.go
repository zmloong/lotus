package ftp_test

import (
	"fmt"
	"testing"

	"log"

	"github.com/zmloong/lotus/sys/ftp"
)

func Test_sys(t *testing.T) {
	sys, err := ftp.NewSys(
		ftp.SetServerIp("118.178.94.244"),
		ftp.SetUsername("userftp"),
		ftp.SetPassword("123456"),
	)
	if err != nil {
		log.Println(err)
	}
	et, err := sys.List("")
	for _, v := range et {
		fmt.Println(v.FileName)
	}
}

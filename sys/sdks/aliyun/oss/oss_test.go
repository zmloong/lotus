package oss_test

import (
	"fmt"
	"testing"

	"github.com/zmloong/lotus/sys/sdks/aliyun/oss"
)

func Test_OSSUploadFile(t *testing.T) {
	sys, err := oss.NewSys(
		oss.SetEndpoint("http://gohitool.oss-accelerate.aliyuncs.com"),
		oss.SetAccessKeyId("xxxxxxx"),
		oss.SetAccessKeySecret("xxxxxxxxxxxxxxxxxxxxxxxxxx"),
		oss.SetBucketName("xxxxxxxxxxx"),
	)
	if err != nil {
		fmt.Printf("初始化OSS 系统失败 err:%v", err)
		t.Logf("初始化OSS 系统失败 err:%s", err.Error())
		return
	} else {
		t.Logf("初始化OSS 成功")
	}
	// if err := CreateBucket("hitoolchat"); err != nil {
	// 	t.Logf("创建 CreateBucket  err:%s", err.Error())
	// } else {
	// 	t.Logf("创建 CreateBucket 成功")
	// }
	if err := sys.UploadFile("test/liwei2dao.jpg", "F:/liwei1dao.jpg"); err != nil {
		t.Logf("上传OSS 系统失败 err:%s", err.Error())
	} else {
		t.Logf("上传OSS 成功")
	}
	// if file, err := GetObject("test/liwei1dao.jpg"); err != nil {
	// 	t.Logf("下载OSS 系统失败 err:%s", err.Error())
	// } else {
	// 	t.Logf("下载OSS 成功 len:%d", len(file))
	// }
}

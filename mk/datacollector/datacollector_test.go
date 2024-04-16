package datacollector

import (
	"fmt"
	"testing"
	"time"
)

/*
 *采集类任务并发模型
 *1.采集器 Runner
 *2.读取器 Reader
 *3.解析器 Parser
 *4.变换器 Transforms
 *5.发送器 Sender
 */
func Test_task(t *testing.T) {
	fmt.Printf("test start...")
	datacollector := NewIRunner()
	datacollector.Init()
	datacollector.Start()
	fmt.Printf("test ing...")
	time.Sleep(20 * time.Second)
	datacollector.Close(3, "手动关闭")
	time.Sleep(2 * time.Second)
	fmt.Printf("test end! \n")
}

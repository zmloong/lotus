package mk

import (
	"fmt"
	"log"
	"sync/atomic"
	"time"
)

const (
	Runner_Stoped   RunnerState = iota //已停止
	Runner_Initing                     //初始化中
	Runner_Starting                    //启动中
	Runner_Runing                      //运行中
	Runner_Stoping                     //关闭中
)

// 采集器实现
func (r *Runner) MaxProcs() int {
	return r.maxProcs // 假设最大处理器数为
}

func (r *Runner) Init() error {
	atomic.StoreInt32(&r.state, int32(Runner_Initing))
	// 初始化逻辑
	r.reader = NewReader(r)
	r.parser = NewParser(r)
	r.transforms = NewTransforms(r)
	r.senders = NewSenders(r)
	return nil
}

func (r *Runner) Start() error {
	atomic.StoreInt32(&r.state, int32(Runner_Starting))
	// 启动逻辑
	r.reader.Start()
	r.parser.Start()
	r.transforms.Start()
	r.senders.Start()
	atomic.StoreInt32(&r.state, int32(Runner_Runing))
	go r.run()

	go func() {
		for {
			if r.state == int32(Runner_Stoping) {
				r.senders.Close()
			}
		}
	}()
	return nil
}

func (r *Runner) Close(state RunnerState, closemsg string) error {

	//初始化启动失败 close
	if r.reader != nil {
		r.reader.Close()
	}

	if state == Runner_Runing {
		log.Println("Runner Close reader succ")
		r.closesignal <- struct{}{}
		close(r.closesignal)
		log.Println("Runner Close closesignal succ")
		close(r.readerPope)
		log.Println("Runner readerPope close")
		close(r.parserPope)
		log.Println("Runner parserPope close")
		close(r.transformsPope)
		log.Println("Runner transformsPope close")
		close(r.sendersPope)
		log.Println("Runner sendersPope close")
		time.Sleep(time.Second) //等待一秒钟 清理剩余数据
	}
	fmt.Println(closemsg)
	// 关闭逻辑
	return nil
}
func (r *Runner) GetreaderPope() chan string {
	return r.readerPope
}
func (r *Runner) GetparserPope() chan string {
	return r.parserPope
}
func (r *Runner) GettransformsPope() chan string {
	return r.transformsPope
}
func (r *Runner) GetsendersPope() chan string {
	return r.sendersPope
}
func (r *Runner) run() {
	defer r.ticker.Stop()
locp:
	for {
		select {
		case <-r.closesignal:
			break locp
		case data := <-r.readerPope:
			log.Println("Runner 开始中转数据")
			r.parserPope <- data
		case <-r.ticker.C:

		}
	}
	log.Println("Runner exit run")
}

// 读取器实现

func (r *Reader) GetRunner() IRunner {
	log.Println("Reader GetRunner")
	return r.runner
}

func (r *Reader) Start() error {
	log.Println("Reader Start")
	go r.collect()
	return nil
}

func (r *Reader) Close() error {
	log.Println("Reader Close")
	return nil
}

func (r *Reader) Input() chan<- string {
	log.Println("Reader Input")
	return r.runner.GetreaderPope()
}
func (r *Reader) collect() chan<- string {
	log.Println("Reader collect")
	for i := 0; i < r.GetRunner().MaxProcs(); i++ {
		go func(i int) {
			for j := 0; j < 10; j++ {
				r.Input() <- fmt.Sprintf("+++ data [%d]-[%d]", i, j)
				log.Println("Reader>>>")
			}
		}(i)
	}

	return r.runner.GetreaderPope()
}

// 解析器实现

func (p *Parser) GetRunner() IRunner {
	return p.runner
}

func (p *Parser) Start() error {
	pipe := p.runner.GetparserPope()
	for i := 0; i < p.runner.MaxProcs(); i++ {
		p.wg.Add(1)
		go p.run(pipe)
	} // 启动逻辑
	return nil
}
func (p *Parser) run(pipe <-chan string) {
	defer p.wg.Done()
	for v := range pipe {
		p.Parse(v)
	}
}
func (p *Parser) Close() error {

	return nil
}

func (p *Parser) Parse(bucket string) {
	log.Println("Reader>>>Parse")
	time.Sleep(time.Second)
	p.runner.GettransformsPope() <- bucket
	log.Println("Parse>>>")
}

// 变换器实现

func (t *Transforms) GetRunner() IRunner {
	return t.runner
}

func (t *Transforms) Start() error {
	pipe := t.runner.GettransformsPope()
	for i := 0; i < t.runner.MaxProcs(); i++ {
		t.wg.Add(1)
		go t.run(pipe)
	}
	return nil
}
func (t *Transforms) run(pipe <-chan string) {
	defer t.wg.Done()
	for v := range pipe {
		t.Trans(v)
	}
}

func (t *Transforms) Close() error {
	// 关闭逻辑
	return nil
}

func (t *Transforms) Trans(bucket string) {
	log.Println("Parse>>>Trans")
	time.Sleep(time.Second)
	t.runner.GetsendersPope() <- bucket
	log.Println("Trans>>>")
}

// 发送器实现
func (s *Sender) GetRunner() IRunner {
	return s.runner
}

func (s *Sender) Start() error {
	pipe := s.runner.GetsendersPope()
	for i := 0; i < s.runner.MaxProcs(); i++ {
		s.Wg.Add(1)
		go s.run(pipe)
	}

	return nil
}

func (s *Sender) Close() error {
	s.runner.Close(Runner_Runing, "end")
	return nil
}

func (s *Sender) Send(bucket string) {
	fmt.Printf("Send out :[%s]\n", bucket)
}
func (s *Sender) run(pipe <-chan string) {
	defer s.Wg.Done()
	for v := range pipe {
		s.Send(v)
	}
}

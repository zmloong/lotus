package datacollector

import (
	"sync"
	"time"
)

type (
	Runner struct {
		Name        string
		maxProcs    int
		state       int32
		closesignal chan struct{}

		reader     IReader
		parser     IParser
		transforms ITransforms
		senders    ISender

		readerPope     chan string
		parserPope     chan string
		transformsPope chan string
		sendersPope    chan string

		ticker *time.Ticker //定时器
	}
	//读取器
	Reader struct {
		Name   string
		runner IRunner
		reader IReader
	}
	//解析器
	Parser struct {
		Name   string
		runner IRunner
		parser IParser
		wg     *sync.WaitGroup
	}
	//变换器
	Transforms struct {
		Name       string
		runner     IRunner
		transforms ITransforms
		wg         *sync.WaitGroup
	}
	//读取器
	Sender struct {
		Name   string
		runner IRunner
		sender ISender
		Wg     *sync.WaitGroup
	}
)

func NewIRunner() (r IRunner) {
	o := &Runner{
		Name:           "Runner",
		maxProcs:       2,
		readerPope:     make(chan string),
		parserPope:     make(chan string),
		transformsPope: make(chan string),
		sendersPope:    make(chan string),
		closesignal:    make(chan struct{}),
		ticker:         time.NewTicker(time.Second),
	}

	return o
}

func NewReader(runner IRunner) (r IReader) {
	o := &Reader{
		runner: runner,
	}
	o.Name = "Reader"
	return o
}

func NewParser(runner IRunner) (p IParser) {
	o := &Parser{
		runner: runner,
		wg:     new(sync.WaitGroup),
	}
	o.Name = "Parser"
	return o
}
func NewTransforms(runner IRunner) (t ITransforms) {
	o := &Transforms{
		runner: runner,
		wg:     new(sync.WaitGroup),
	}
	o.Name = "Transforms"
	return o
}
func NewSenders(runner IRunner) (s ISender) {
	o := &Sender{
		runner: runner,
		Wg:     new(sync.WaitGroup),
	}
	o.Name = "Sender"
	return o
}

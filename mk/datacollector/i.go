package datacollector

type (
	RunnerState int32
	//采集器
	IRunner interface {
		MaxProcs() (maxprocs int)
		Init() (err error)
		Start() (err error)
		Close(state RunnerState, closemsg string) (err error)

		GetreaderPope() chan string
		GetparserPope() chan string
		GettransformsPope() chan string
		GetsendersPope() chan string
	}
	//读取器
	IReader interface {
		GetRunner() IRunner
		Start() (err error)
		Close() (err error)
		Input() chan<- string
	}
	//解析器
	IParser interface {
		GetRunner() IRunner
		Start() (err error)
		Close() (err error)
		Parse(bucket string)
	}
	//变换器
	ITransforms interface {
		GetRunner() IRunner
		Start() (err error)
		Close() (err error)
		Trans(bucket string)
	}
	//读取器
	ISender interface {
		GetRunner() IRunner
		Start() (err error)
		Close() (err error)
		Send(bucket string)
	}
)

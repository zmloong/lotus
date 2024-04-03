package redis

import (
	"errors"
	"time"
)

type RedisMutex struct {
	sys    IRedis
	key    string
	expiry int //过期时间 单位秒
	delay  time.Duration
}

func (this *Redis) NewRedisMutex(key string, optfns ...RMutexOptionfn) (result *RedisMutex, err error) {
	opt := newRMutexOptions(optfns...)
	result = &RedisMutex{
		sys:    this.client,
		key:    key,
		expiry: opt.expiry,
		delay:  opt.delay,
	}
	return
}

func (this *RedisMutex) Lock() (err error) {

	wait := make(chan error)

	go func() {
		start := time.Now()

		for int(time.Now().Sub(start).Seconds()) < this.expiry {
			if res, err := this.sys.Lock(this.key, this.expiry); err == nil && res {
				wait <- nil
				return
			} else if !res {
				time.Sleep(this.delay)
			} else {
				wait <- err
			}
		}
		wait <- errors.New("time out")
	}()

	wait <- err
	return
}
func (this *RedisMutex) Unlock() {
	this.sys.UnLock(this.key)
}

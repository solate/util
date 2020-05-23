package retry

import (
	"errors"
	"time"
)

var errMaxRetriesReached = errors.New("exceeded retry limit")

// 重复调用方法. 方法返回nil时停止
// fn: 重试方法
// n: 重试次数
// s: 重试间隔时间
func Do(fn func() error, n int, t time.Duration) (err error) {

	if err := fn(); err != nil {
		if n--; n > 0 {
			time.Sleep(t)
			return Do(fn, n, t)
		} else {
			return errMaxRetriesReached
		}
	}

	return
}

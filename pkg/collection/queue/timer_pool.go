package queue

import (
	"sync"
	"time"
)

type TimerPool interface {
	Get(time.Duration) *time.Timer
	Put(*time.Timer)
}

type timerPool struct {
	timerPool sync.Pool
}

func (a *timerPool) Get(timeout time.Duration) *time.Timer {
	t := a.timerPool.Get().(*time.Timer)
	t.Reset(timeout)
	return t
}

func (a *timerPool) Put(t *time.Timer) {
	t.Stop()
	a.timerPool.Put(t)
}

var TimerRecycler = NewTickerPool()

func NewTickerPool() TimerPool {
	return &timerPool{
		timerPool: sync.Pool{
			New: func() interface{} {
				t := time.NewTimer(time.Hour)
				t.Stop()
				return t
			},
		},
	}
}

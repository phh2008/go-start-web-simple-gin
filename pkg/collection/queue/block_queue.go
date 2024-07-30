package queue

import (
	"fmt"
	"sync/atomic"
	"time"
)

type BlockQueue[T any] interface {
	// Put 添加元素，如果队列已满，会阻塞直到添加成功
	Put(element T) bool

	// Offer 添加元素，如果队列已满，会等待直到超时
	Offer(element T, timeout time.Duration) bool

	// Take 获取元素并移除元素，如果队列为空，会阻塞直到获取成功
	Take() T

	// Poll 获取元素并移除元素，如果队列为空，会等待直到超时
	Poll(timeout time.Duration) (T, bool)

	// Len 队列长度
	Len() uint32

	// Cap 队列容量
	Cap() uint32
}

type blockingQueue[T any] struct {
	buffer    chan T
	capacity  uint32
	len       uint32
	timerPool TimerPool
}

// NewBlockingQueue 创建一个阻塞队列
// capacity 队列容量,必须大于0，否则 panic
func NewBlockingQueue[T any](capacity uint32) BlockQueue[T] {
	if capacity <= 0 {
		panic(fmt.Sprintf("capacity illegal: %d, must be greater than 0", capacity))
	}
	return &blockingQueue[T]{
		capacity:  capacity,
		buffer:    make(chan T, capacity),
		timerPool: TimerRecycler, // a pool to reuse timer
		len:       0,
	}
}

func (a *blockingQueue[T]) Offer(element T, timeout time.Duration) bool {
	t := a.timerPool.Get(timeout)
	select {
	case <-t.C:
		a.timerPool.Put(t)
		return false
	case a.buffer <- element:
		atomic.AddUint32(&a.len, 1)
		a.timerPool.Put(t)
		return true
	}
}

func (a *blockingQueue[T]) Put(element T) bool {
	a.buffer <- element
	atomic.AddUint32(&a.len, 1)
	return true
}

func (a *blockingQueue[T]) Take() T {
	val := <-a.buffer
	atomic.AddUint32(&a.len, ^uint32(1-1))
	return val
}

func (a *blockingQueue[T]) Poll(timeout time.Duration) (T, bool) {
	t := a.timerPool.Get(timeout)
	select {
	case <-t.C:
		// timeout, return false
		var t T
		return t, false
	case val := <-a.buffer:
		atomic.AddUint32(&a.len, ^uint32(1-1))
		a.timerPool.Put(t)
		return val, true
	}
}

func (a *blockingQueue[T]) Len() uint32 {
	return a.len
}

func (a *blockingQueue[T]) Cap() uint32 {
	return a.capacity
}

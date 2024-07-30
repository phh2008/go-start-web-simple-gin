package queue

import (
	"fmt"
	"testing"
	"time"
)

func TestBlockQueueTimeout(t *testing.T) {
	mq := NewBlockingQueue[string](10)
	go func() {
		for {
			// 超时获取
			time.Sleep(time.Millisecond * 500)
			v, ok := mq.Poll(time.Millisecond * 900)
			if ok {
				t.Log("poll  success:", time.Now().Format("2006-01-02 15:04:05.000"), mq.Len(), mq.Cap(), v)
			} else {
				t.Log("poll  timeout:", time.Now().Format("2006-01-02 15:04:05.000"), mq.Len(), mq.Cap())
			}
		}
	}()
	tk := time.NewTicker(time.Second * 1)
	for i := 0; i < 5; i++ {
		go func(i int) {
			for {
				// 定时添加元素
				//<-tk.C
				time.Sleep(time.Millisecond * 50)
				ok := mq.Offer(fmt.Sprintf("hello-%d", i), time.Millisecond*200)
				if ok {
					t.Log("offer success:", time.Now().Format("2006-01-02 15:04:05.000"), mq.Len(), mq.Cap())
				} else {
					t.Log("offer timeout:", time.Now().Format("2006-01-02 15:04:05.000"), mq.Len(), mq.Cap())
				}
			}
		}(i)
	}

	<-time.After(time.Second * 30)
	tk.Stop()
}

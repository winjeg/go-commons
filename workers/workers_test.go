package workers

import (
	"errors"
	"testing"
	"time"
)

func TestNewPool(t *testing.T) {
	// 创建一个包含 3 个 Worker 的线程池
	pool := NewPool("demo", 3)
	pool.Start()

	// 添加 10 个任务到线程池
	for i := 0; i < 30; i++ {
		pool.AddJob(Job{ID: i + 1, Task: doDemoWork})
	}
	// 停止线程池
	pool.Stop()
}

func doDemoWork() error {
	time.Sleep(time.Millisecond * 100)
	return errors.New("error occcured")
}

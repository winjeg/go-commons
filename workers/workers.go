package workers

import (
	"github.com/winjeg/go-commons/exception"
	"github.com/winjeg/go-commons/log"

	"fmt"
	"sync"
)

// Job 结构体表示要执行的任务
type Job struct {
	ID   int
	Task func() error
}

// Worker 结构体表示工作线程
type Worker struct {
	ID         int
	Name       string
	JobChannel chan Job
	Quit       chan bool
}

// NewWorker 创建一个新的 Worker
func NewWorker(id int, jobChannel chan Job, name string) *Worker {
	return &Worker{
		ID:         id,
		JobChannel: jobChannel,
		Quit:       make(chan bool),
		Name:       name,
	}
}

// Start 启动 Worker
func (w *Worker) Start(wg *sync.WaitGroup) {
	go func() {
		defer exception.Catch()
		defer wg.Done()
		for {
			select {
			case job := <-w.JobChannel:
				if err := job.Task(); err != nil {
					log.GetLogger(nil).Errorf("error running  worker: %s job id: %d, err: %+v", w.Name, job.ID, err)
				}
			case <-w.Quit:
				log.GetLogger(nil).Infof("Worker %s stopped\n", w.Name)
				return
			}
		}
	}()
}

// Stop 停止 Worker
func (w *Worker) Stop() {
	go func() {
		w.Quit <- true
	}()
}

// Pool 结构体表示线程池
type Pool struct {
	Workers    []*Worker
	JobChannel chan Job
	wg         sync.WaitGroup
}

// NewPool 创建一个新的线程池
func NewPool(name string, numWorkers int) *Pool {
	jobChannel := make(chan Job)
	workers := make([]*Worker, numWorkers)

	for i := 0; i < numWorkers; i++ {
		workerName := fmt.Sprintf("%s-%d", name, i)
		workers[i] = NewWorker(i, jobChannel, workerName)
	}

	return &Pool{
		Workers:    workers,
		JobChannel: jobChannel,
	}
}

// Start 启动线程池
func (p *Pool) Start() {
	for _, worker := range p.Workers {
		p.wg.Add(1)
		worker.Start(&p.wg)
	}
}

// Stop 停止线程池
func (p *Pool) Stop() {
	for _, worker := range p.Workers {
		worker.Stop()
	}
	p.wg.Wait()
}

// AddJob 添加任务到线程池
func (p *Pool) AddJob(job Job) {
	p.JobChannel <- job
}

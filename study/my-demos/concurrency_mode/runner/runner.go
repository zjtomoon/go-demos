package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

var (
	ErrTimeout  = errors.New("cannot finish tasks within the timeout")
	ErrInterrpt = errors.New("received interrupt from OS")
)

// Runner给定一系列的Task,要求在规定的timeout内跑完，不然就报错
//如果操作系统给了中断信号，也报错

type Runner struct {
	interrupt chan os.Signal
	complete  chan error
	timeout   <-chan time.Time //用来计时
	tasks     []func(int)      //task的列表
}

func New(t time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1), //TODO: why buffer is 1: select 操作带缓存的channel
		complete:  make(chan error),
		timeout:   time.After(t),
		tasks:     make([]func(int), 0),
	}
}

func (r *Runner) AddTask(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		select {
		case <-r.interrupt:
			//说明操作系统传递了interrupt的信号
			signal.Stop(r.interrupt)
			return ErrInterrpt
		default:
			task(id)
		}
	}
	return nil
}

func (r *Runner) Start() error {
	//relay interrupt from OS
	signal.Notify(r.interrupt, os.Interrupt)

	//Run the task
	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

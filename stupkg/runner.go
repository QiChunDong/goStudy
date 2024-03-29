package stupkg

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"time"
)

// 在给定的超时时间内执行一组代码
// 在操作系统发送中断信号时结束任务
type Runner struct {
	// 通道报告
	// 从操作系统发送的信号 如 ctrl+c
	interrupt chan os.Signal
	// 报告任务已经处理完成
	complete chan error
	// 报告任务已经超时
	timeout <-chan time.Time
	// 一组已索引顺序依次执行的函数
	tasks []func(int)
}

// 超时时会返回该异常
var ErrorTimeout = errors.New("received timeout")
// 操作系统中断事件返回该异常
var ErrorInterrupt = errors.New("received interrupt")

// 新的准备使用的Runner 指针
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete: make(chan error),
		timeout: time.After(d),
	}
}

// 将任务附加到Runner上，接受int参数的函数
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks= append(r.tasks, tasks...)
}

// 执行任务 返回异常 可能是中断也可能是超时
func (r *Runner) Start() error {
	// 我们希望收到所有的终端信号
	signal.Notify(r.interrupt, os.Interrupt)

	// 并发执行
	go func() {
		r.complete <- r.run()
	}()
	
	select {
		case err := <-r.complete:
			return err
		case <-r.timeout:
			return ErrorTimeout
	}
}

// 按顺序执行任务
func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.gotTnterrupt() {
			return ErrorTimeout
		}
		task(id)
	}
	return nil
}

func (r *Runner) gotTnterrupt() bool {
	select {
	case <-r.interrupt:
		// 向信号组件发送接收到的任何中断信号
		// 自杀
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}

}

// 超时时间 3秒
const timeout = 3 * time.Second

func TestRunner() {
	fmt.Println("Start Work")

	r := New(timeout)

	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err{
		case ErrorTimeout:
			fmt.Println("Terminating due to timeout.")
			os.Exit(1)
		case ErrorInterrupt:
			fmt.Println("Terminating due to interrupt.")
			os.Exit(2)
		}
	}

	fmt.Println("process ended.")
}

func createTask() func(int) {
	return func(id int) {
		fmt.Printf("Processor - Task #%d\n", id)
		time.Sleep(time.Duration(id-1) * time.Second)
	}
}
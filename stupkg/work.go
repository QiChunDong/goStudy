// 管理一个goroutine池来完成工作
// 非阻塞的执行任务
package stupkg

import (
	"log"
	"sync"
	"time"
)

// 实现了Worker的接口  才能使用池
type Worker interface {
	Task()
}

// goroutine池
// 完成任何已提交的worker
type WPool struct {
	work chan Worker
	wg sync.WaitGroup
}

// 新的池
func NewWorkerPool(maxgoroutines int) *WPool {
	p := WPool{
		work: make(chan Worker),
	}

	p.wg.Add(maxgoroutines)

	for i := 0; i < maxgoroutines; i++ {
		go func(){
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}
	return &p
}

// 执行  向池里的通道写入work
func (p *WPool) Run(w Worker) {
	p.work <-w
}

func (p *WPool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}

var names = []string{
	"steve",
	"emily",
	"john",
	"hh",
	"mask",
}

type NamePrinter struct{
	name string
}


func (pr *NamePrinter) Task() {
	log.Println(pr.name)
	time.Sleep(time.Second)
}


func TestWorker() {
	p := NewWorkerPool(2)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++{
		for _, name := range names {
			nm := NamePrinter{
				name: name,
			}
			// 协程执行任务
			// Run 结束表示执行结束
			go func() {
				p.Run(&nm)
				wg.Done()
			}()
		}
	}

	wg.Wait()
	// 工作完成  停止工作
	p.Shutdown()
}

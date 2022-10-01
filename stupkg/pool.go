package stupkg

import (
	"errors"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// 管理可以在多个gorouine之间共享的资源
// 必须要实现io.Closer()
type Pool struct {
	m sync.Mutex // 资源访问要加互斥锁
	resources chan io.Closer
	factory func() (io.Closer, error)
	closed bool
}

// 请求了一个已经关闭的pool
var ErrorPoolClosed = errors.New("Pool has been closed")

func NewPool(fn func() (io.Closer, error), size int) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size value too small")
	}

	return &Pool{
		factory: fn,
		resources: make(chan io.Closer, size),
	}, nil
}

// 获取一个资源
func (p *Pool) Aquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		log.Println("Aquire:", "Shared resource")
		if !ok {
			return nil, ErrorPoolClosed
		}
		return r, nil
	default:
		log.Println("Aquire:", "New resource")
		return p.factory()
	}
}

func (p *Pool) Release(r io.Closer) {
	// 加互斥锁
	p.m.Lock()
	defer p.m.Unlock()

	// 池都关了  直接关闭资源
	if p.closed {
		r.Close()
		return
	}

	select {
	case p.resources <- r:
		log.Println("Relase:", "In Queue")
	default:
		log.Println("Relase:", "closing")
		r.Close()
	}
}

func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}

	p.closed = true

	// 这是通道的关闭
	close(p.resources)

	// 挨个关闭资源
	for r := range p.resources {
		r.Close()
	}

}

const (
	maxGoroutines = 25 // 协程数
	polledResources = 5 // 池中的资源数
)

type dbConnection struct {
	ID int32
}

// 关闭链接
func (dbConn *dbConnection) Close() error {
	log.Println("Close: Connection", dbConn.ID)
	return nil
}

// 分配独一无二的id
var idCounter int32

// 创建一个连接  pool的参数
func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)
	return &dbConnection{id}, nil
}

func TestPool() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	p, err := NewPool(createConnection, polledResources)
	if err != nil {
		log.Println(err)
	}

	for query := 0; query < maxGoroutines; query++ {
		go func (q int)  {
			performQuries(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()
	log.Println("Shutdown program")
	p.Close()
}


// 测试连接过程
func performQuries(query int, p *Pool){
	conn, err := p.Aquire()

	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close()

	// 模拟等待
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d] cid[%d]\n", query, conn.(*dbConnection).ID)
}

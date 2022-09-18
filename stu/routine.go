package stu

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func TestRoutine() {
	// 分配处理器给调度器  1和多个的区别是并发与并行
	runtime.GOMAXPROCS(2)

	// 计数器 表示等待2个rotine执行
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Gorotines")
	go func() {
		// defer 类似于finally 惩处推出前总会执行
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()
	go func() {
		// defer 类似于finally 惩处推出前总会执行
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("\nFinished!!!!")
}

var wg sync.WaitGroup

func TestPrime() {
	runtime.GOMAXPROCS(1)
	wg.Add(2)
	fmt.Println("Create two goroutins")
	go printPrime("A")
	go printPrime("B")
	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("\nFinished!!!!")
}

// 显示5000以内的素数
func printPrime(prefix string) {
	// 通知完成
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s: %d\n", prefix, outer)
	}
	fmt.Printf("%s Completed", prefix)
}

// 测试协程竞争
var (
	counter int64
	wg2     sync.WaitGroup
)

// 测试协程竞争
func TestComp() {
	wg2.Add(2)

	// 有竞争
	go incCounter(1)
	go incCounter(2)

	wg2.Wait()
	fmt.Println("Final Counter: ", counter)
}

// 测试协程竞争
func TestCompAtom() {
	wg2.Add(2)

	// 加锁的
	go incCounterAtom(1)
	go incCounterAtom(2)

	wg2.Wait()
	fmt.Println("Final Counter: ", counter)
}

func incCounter(id int) {
	defer wg2.Done()

	for count := 0; count < 2; count++ {
		// 捕获counter
		value := counter

		// 当前goroutine 从线程退出  并且放回队列
		runtime.Gosched()

		value++
		// 保存给counter
		counter = value
	}
}

// 加原子锁的
func incCounterAtom(id int) {
	defer wg2.Done()

	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}
}

var (
	shutDown int64
	wg3      sync.WaitGroup
)

func TestAtom() {
	wg3.Add(2)

	go doWork("A")
	go doWork("B")

	// 等待一段时间  等待
	time.Sleep(1 * time.Second)

	// 这是应该停止  将shutdown 置为1
	fmt.Println("Shutdown now!")
	// 原子修改值
	atomic.StoreInt64(&shutDown, 1)

}

func doWork(name string) {
	defer wg3.Done()

	for {
		fmt.Printf("Doing %s Work \n", name)
		time.Sleep(250 * time.Millisecond)

		// 原子加载值
		// 只有原子函数判断到 shutDown==1  停止
		if atomic.LoadInt64(&shutDown) == 1 {
			fmt.Printf("Shutting down %s \n", name)
			break
		}
	}
}

var (
	counterMutex int
	wg4          sync.WaitGroup
	mutex        sync.Mutex
)

func TestMutex() {
	wg4.Add(2)
	go incCounterMutex(1)
	go incCounterMutex(2)
	wg4.Wait()
	fmt.Println("Final counterMutex: ", counterMutex)

}

// 测试互斥锁 mutex
// mutual exclusion 互斥的概念
// 创建临界区 只能有一个goroutine访问
func incCounterMutex(id int) {
	defer wg4.Done()

	for count := 0; count < 2; count++ {
		// 加锁，同一时刻只允许一个goroutine进入
		mutex.Lock()
		// 以下就是加了互斥锁的临界区
		{
			// 捕获counterMutex
			value := counterMutex

			// 当前goroutine 从线程退出  并且放回队列
			runtime.Gosched()

			value++
			// 保存给counterMutex
			counterMutex = value
		}
		mutex.Unlock()
		// 释放锁
		// 其他等待的goroutine就可以进入临界区了
	}
}

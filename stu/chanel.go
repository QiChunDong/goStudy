package stu

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	wg_c1 sync.WaitGroup
	wg_c2 sync.WaitGroup
	wg_c3 sync.WaitGroup
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 测试无缓冲的通道
func TestNoBufferChanel() {
	wg_c1.Add(2)

	court := make(chan int)

	go player("Tom", court)
	go player("John", court)

	// 初始值 1
	court <- 1
	wg.Wait()

	fmt.Println("Game Over!")
}

func player(name string, court chan int) {
	defer wg_c1.Done()

	for {
		ball, ok := <-court
		if !ok {
			// 通道被关闭  结束 赢了
			fmt.Printf("Player %s won !\n", name)
			return
		}

		// 选随机数 来判断是否发球
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s missed ! \n", name)

			// 关闭通道  结束 输了
			close(court)
			return
		}

		fmt.Printf("Player %s Hit %d \n", name, ball)
		ball++

		// 通道塞值  球打回对手
		court <- ball
	}
}

// 测试无缓冲的通道案例2
func TestNoBufferChanel2() {
	wg_c2.Add(1)
	baton := make(chan int)

	fmt.Println("Run Start!")
	go run(baton)

	// 比赛开始
	baton <- 1

	// 接力完成了
	wg_c2.Wait()
	fmt.Println("Run End!")
}

// 案例: 接力跑
// 跑者跑一圈接力给下一位
// 如果达到4圈 停
func run(baton chan int) {
	// 预定义下一个运动员
	var newRunner int
	// 等待接力棒
	runner := <-baton

	fmt.Printf("Runner %d is running!\n", runner)

	// 跑起来的同时 创建 运动员
	// 下一位等待入列
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d to the Line\n", newRunner)
		go run(baton)
	}

	// 跑起来
	time.Sleep(100 * time.Microsecond)

	if runner == 4 {
		fmt.Printf("Runner %d finished!\n", newRunner)
		wg_c2.Done()
		return
	}

	// 接力
	fmt.Printf("Runner %d exchanged by %d\n", runner, newRunner)

	baton <- newRunner
}

const (
	numberGroroutines = 4
	taskLoad          = 10
)

// 测试有缓冲的通道
func TestBufferChanel() {
	wg_c3.Add(numberGroroutines)
	// 所有任务的通道
	tasks := make(chan string, taskLoad)

	for gr := 1; gr <= numberGroroutines; gr++ {
		go work(tasks, gr)
	}

	// 投递任务
	// 向通道写数据
	fmt.Println("Task Start !")
	for t := 0; t < taskLoad; t++ {
		tasks <- fmt.Sprintf("Task %d", t)
	}

	// 写结束了 关闭通道
	close(tasks)

	wg_c3.Wait()
	fmt.Println("All Tasks Done")
}

func work(tasks chan string, worker int) {
	defer wg_c3.Done()

	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("Worker %d, shut down! \n", worker)
			return
		}

		fmt.Printf("Worker %d, Start work: %s !\n", worker, task)

		// 处理task
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Printf("Worker %d, end work:  %s !\n", worker, task)
	}
}

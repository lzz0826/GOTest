package code

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func ContextTest() {

	//创建上下文
	ctx := context.Background()

	// 设置超时控制，3 秒后超时
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel() // 延迟取消上下文，确保及时释放资源cancel

	var wg sync.WaitGroup
	wg.Add(1)

	//创建结果通道
	resultChan := make(chan string)

	//开另一个线程会ContextTest 函数会立即返回 不会等待 doSome 函数执行完成 会导致doSome无法执行完
	go doSome(ctx, 2*time.Second, &wg, resultChan)

	select {
	case some := <-resultChan: //接收 渠道
		fmt.Println("Received result:", some)
	case <-time.After(10 * time.Second): // 超时处理
		fmt.Println("ContextTest -- Timeout occurred")
	}

	//看情况去开另一个现成 去等待(有可能ContextTest已经先结束了 wg.Wait()执行到
	go func() {
		wg.Wait()
		close(resultChan) // 关闭通道
	}()
	fmt.Println("ContextTest End")

}

// resultChan chan<- string 返回结果渠道
func doSome(ctx context.Context, duration time.Duration, wg *sync.WaitGroup, resultChan chan<- string) {
	defer wg.Done()
	select {
	case <-ctx.Done():
		fmt.Println("Operation timed out or canceled:", ctx.Err())
		resultChan <- "" // 发送空字符串到结果通道
	case <-time.After(duration):
		fmt.Println("Operation completed successfully")
		//另外模拟ContextTest 等待 doSome 超时
		//time.Sleep(11 * time.Second)
		resultChan <- "Some..." // 发送空字符串到结果通道
	}
}

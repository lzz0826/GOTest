package channelTest

import (
	"fmt"
	"sync"
)

// -----1. 单向通道（unidirectional channel）：-----
func UnidirectionalSend(ch chan<- string, str string) {
	ch <- str
	//close(ch)
}
func UnidirectionalSendReceive(ch chan string) {
	for value := range ch {
		fmt.Printf(value)
	}
}

// -----2. 带缓冲的通道（buffered channel）：-----
var ch chan string

func InitChan() {
	//带缓冲的通道 缓冲满了新数据会在外面等待
	ch = make(chan string, 5)
}

func BufferedChannel() {
	ch <- "1"
	ch <- "2"
	ch <- "3"
	ch <- "4"
	ch <- "5"
	ch <- "6"
}

func PChan() {
	// range 语句会阻塞等待，直到通道被关闭。
	//for s := range ch {
	//	fmt.Println(s)
	//}
	for i := 0; i < 3; i++ {
		data := <-ch
		fmt.Println(data)
	}
}

// -----3. 使用 select 语句处理多个通道：-----

var ch1 chan string
var ch2 chan string

func InitCh1Ch2() {
	ch1 = make(chan string)
	ch2 = make(chan string)
}

func Ch1() {
	ch1 <- "Ch1Data"
}

func CloseC1() {
	close(ch1)
}

func Ch2() {
	ch2 <- "Ch2Data"
}

func SelectCh1Ch2() {
	for {
		select {
		case data1, ok := <-ch1:
			//监测通道有无被关闭
			if !ok {
				fmt.Println("ch1 closed")
				return
			}
			fmt.Println(data1)
		case data2, ok := <-ch2:
			if !ok {
				fmt.Println("ch2 closed")
				return
			}
			fmt.Println(data2)
		}
	}
}

// -----4. 使用通道实现同步：-----

//syncLock.WaitGroup 类似 JAVA JUC CountDownLatch

func Worker(id int, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done() // 在函数结束时减少计数器（defer）函数的执行直到当前函数返回之前。
	ch <- fmt.Sprintf("Worker %d done", id)
}

func ChanWorker() {
	ch := make(chan string)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Worker(i, &wg, ch) // 使用goroutine调用Worker函数
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for value := range ch {
		fmt.Println(value)
	}

}

package syncLock

import (
	"fmt"
	"sync"
	"time"
)

//syncLock.Cond：是一个条件变量，用于在多个 goroutine 之间同步操作。 类似java.util.concurrent.locks.Condition
//条件变量允许一个或多个 goroutine 等待或通知事件的发生。使用条件变量时，goroutine
//可以等待某个条件的发生，或者在条件满足时通知其他等待的 goroutine。条件变量必须与互斥锁一起使用，以保护共享变量的访问。

//调用 runtime.notifyListAdd 将等待计数器加一并解锁
//调用 runtime.notifyListWait 等待其他 Goroutine 的唤醒并加锁

//sync.Cond.Signal 方法会唤醒队列最前面的 Goroutine
//sync.Cond.Broadcast 方法会唤醒队列中全部的 Goroutine

func DoCond() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	done := make(chan string, 2)

	// 第一个goroutine等待条件满足

	go func() {
		id := "123"
		defer func(st string) {
			done <- "线程id : " + st
		}(id)
		mu.Lock()         // 上锁
		defer mu.Unlock() // 在函数结束时解锁，这里的 defer 会在函数结束前执行，所以会先解锁
		fmt.Println("等待条件...")
		cond.Wait() // 等待条件满足，这里会释放互斥锁并阻塞当前 goroutine，等待 cond.Signal() 的通知
		fmt.Println("条件满足，继续执行")
	}()

	// 第二个goroutine触发条件满足
	go func() {
		id := "456"
		defer func(st string) {
			done <- "线程id : " + st
		}(id)
		time.Sleep(time.Second * 6) // 2秒后触发条件
		mu.Lock()                   // 上锁
		defer mu.Unlock()           // 在函数结束时解锁
		fmt.Println("条件已满足")
		cond.Signal() // 触发条件满足，通知等待的 goroutine，这里不会释放互斥锁
	}()

	// 等待两个goroutine完成
	//<-done // 等待第一个 goroutine 完成
	//<-done // 等待第二个 goroutine 完成

	go func() {

		for s := range done {
			fmt.Println(s)
		}
	}()

}

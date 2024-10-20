package syncLock

import (
	"sync"
	"time"
)

//syncLock.Mutex：
//是一个互斥锁，用于在代码块上加锁和解锁，以保护共享资源不被多个 goroutine 同时访问。
//一次只能有一个 goroutine 持有 Mutex 的锁，其他 goroutine 在获取锁之前会被阻塞。
//Mutex 使用较为简单，但如果有大量的读操作和少量的写操作时，性能可能较低。

var (
	count int
	mu    sync.Mutex
)

func AddCount(i int) {
	mu.Lock()
	time.Sleep(4 * time.Second)
	count = count + i
	defer mu.Unlock()
}

func GetCount() int {
	mu.Lock()
	defer mu.Unlock()

	return count
}

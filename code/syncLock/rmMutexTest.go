package syncLock

import (
	"sync"
)

//syncLock.RWMutex： 读写锁
//是一个读写互斥锁，比 Mutex 更加灵活。它允许多个 goroutine 并发地读取共享资源，但在有写操作时需要排他性访问。
//当有 goroutine 持有写锁时，其他所有 goroutine 都被阻塞，包括读操作。
//但当有 goroutine 持有读锁时，其他 goroutine 也可以继续持有读锁，从而实现了并发读取。

type RmMutex struct {
	Count int
	rm    sync.RWMutex
}

func (rm *RmMutex) AddCount(i int) {
	rm.rm.Lock()
	rm.Count = rm.Count + i
	defer rm.rm.Unlock()
}

func (rm *RmMutex) GetCount() int {
	rm.rm.Lock()

	defer rm.rm.Unlock()
	return rm.Count
}

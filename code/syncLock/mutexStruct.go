package syncLock

import (
	"sync"
	"time"
)

//syncLock.Mutex：
//是一个互斥锁，用于在代码块上加锁和解锁，以保护共享资源不被多个 goroutine 同时访问。
//一次只能有一个 goroutine 持有 Mutex 的锁，其他 goroutine 在获取锁之前会被阻塞。
//Mutex 使用较为简单，但如果有大量的读操作和少量的写操作时，性能可能较低。

// 以构造方式
type MutexStruct struct {
	count int
	//锁
	mu sync.Mutex
}

func (ms *MutexStruct) CountAdd() {

	//加锁
	ms.mu.Lock()

	time.Sleep(4 * time.Second)
	ms.count++
	//解所
	defer ms.mu.Unlock()
}

func (ms *MutexStruct) GetCount() int {

	ms.mu.Lock()

	defer ms.mu.Unlock()

	return ms.count

}

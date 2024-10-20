package syncLock

import (
	"fmt"
	"sync"
)

func OnceTest() {
	var once sync.Once

	for i := 0; i < 5; i++ {
		go func() {
			once.Do(doSomeThing)
			//doSomeThing()
		}()
	}

}

func doSomeThing() {
	fmt.Printf("DoSomeThing....")
}

//重制 Once

var once2 sync.Once

func DoSomeThingR() {
	fmt.Printf("R....")
	SetOnce(doSomeThing)
}

func SetOnce(fun func()) {
	once2.Do(fun)
}

func ResetOnce() {
	once2 = sync.Once{}
}

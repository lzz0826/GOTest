package code

import (
	"fmt"
	"strconv"
	"sync"
)

// 迭代数组或切片：
func RangeArrAndSlice() {
	arr := [4]int{1, 2, 3, 4}
	for index, value := range arr {
		fmt.Printf("index is : %d value is : %d\n", index, value)
	}

	slice := []int{1, 2, 3, 4}
	for index, value := range slice {
		fmt.Printf("index is : %d value is : %d\n", index, value)
	}

}

// 迭代映射（map）：
func RangeArrMap() {
	//m := make(map[string]string)
	//m["1"] = "ton"
	m := map[string]string{"1": "ton", "2": "ed"}
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}

func RangeChar() {
	str := "hi my is"
	for index, char := range str {
		fmt.Printf("Index: %d, Char: %c\n", index, char)
	}
}

func RangeChan() {
	ch := make(chan string)

	var wg sync.WaitGroup // 创建 WaitGroup 对象 类似Java CountDownLatch

	for i := 0; i < 10; i++ {
		wg.Add(1) // 增加等待的goroutine数量
		go func(i int) {
			defer wg.Done() // 减少等待的goroutine数量
			ch <- strconv.Itoa(i)
		}(i) // 传递 i 的值作为参数
	}

	go func() {
		wg.Wait() // 等待所有的goroutine完成
		close(ch) // 关闭通道
	}()

	for value := range ch {
		fmt.Println(value)
	}

}

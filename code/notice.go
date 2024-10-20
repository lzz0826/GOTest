package code

import "fmt"

func Notice() {
	// := **不是修改副本特例 坑

	//**正确使用副本
	// 创建一个原始分数切片的副本
	//o := make([]int, len(score))
	//copy(o, score)

	//---------------切片---------------
	// 原始切片
	originalSlice := []int{1, 2, 3}

	// 创建一个新的切片变量，但没有使用 & 符号
	newSlice := originalSlice

	// 修改新切片的第一个元素
	newSlice[0] = 100

	fmt.Println("原始切片：", originalSlice) // 输出: [100 2 3]
	fmt.Println("新切片：", newSlice)       // 输出: [100 2 3]

	//---------------map-------------------
	// 原始映射
	originalMap := map[string]int{"a": 1, "b": 2}

	// 创建一个新的映射变量，值为原始映射的副本
	newMap := originalMap

	// 修改新映射的值
	newMap["a"] = 100

	fmt.Println("原始映射：", originalMap) // 输出: map[a:100 b:2]
	fmt.Println("新映射：", newMap)       // 输出: map[a:100 b:2]

	//-----------------通道-----------------

	// 原始通道
	originalChannel := make(chan int)

	// 创建一个新的通道变量，值为原始通道的副本
	newChannel := originalChannel

	// 向新通道发送值
	go func() {
		newChannel <- 1
	}()

	// 从原始通道接收值
	value := <-originalChannel

	fmt.Println("从原始通道接收到的值：", value) // 输出: 1

	//-----------数组(要用&才是修改远本的 := 是副本)------------------
	// 原本数组
	originalArray := [3]int{1, 2, 3}

	// 指向原本的指针
	pointToOriginal := &originalArray

	// 修改指向原本的数组
	pointToOriginal[0] = 100

	fmt.Println("原本数组：", originalArray)       // 输出: [100 2 3]
	fmt.Println("指向原本的数组：", *pointToOriginal) // 输出: [100 2 3]

	// 创建副本
	copyOfArray := originalArray

	// 修改副本数组
	copyOfArray[1] = 200

	fmt.Println("原本数组：", originalArray) // 输出: [100 2 3]
	fmt.Println("副本数组：", copyOfArray)   // 输出: [100 200 3]

}

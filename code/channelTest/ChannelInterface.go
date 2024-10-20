package channelTest

import "fmt"

type Person struct {
	Name string
	Age  int
}

// 定义一个通用的通道类型，用来传送任意类型的数据
//type GenericChannel chan interface{}

func SetChat(ch chan interface{}) {
	go func() {
		ch <- "ts"
		ch <- 1
		ch <- Person{Name: "n1", Age: 12}
		close(ch)
	}()
}

func ChatType(ch chan interface{}) {
	// 从通道接收数据
	for data := range ch {
		switch v := data.(type) {
		case int:
			fmt.Println("Received an integer:", v)
		case string:
			fmt.Println("Received a string:", v)
		case Person:
			fmt.Println("Received a person:", v.Name, v.Age)
		default:
			fmt.Println("Received an unknown type")
		}
	}
}

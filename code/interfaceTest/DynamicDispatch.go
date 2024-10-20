package interfaceTest

import "fmt"

// 动态派发 多态
type Duck interface {
	Quack()
	Jup()
}

// Cat实现
type Cat struct {
	Name string
}

type Dog struct {
	Name string
	Age  int
}

func (d Dog) Jup() string {
	fmt.Printf(d.Name + "Is Jup")
	return d.Name
}

func (d Dog) Quack() string {
	fmt.Printf(d.Name + "Is JUP")
	return d.Name
}

// (c Cat) 是方法的接收者(receiver)
func (c Cat) Quack() string {
	fmt.Printf(c.Name + "Is Quack")
	return c.Name
}

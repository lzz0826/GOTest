package code

import (
	"fmt"
)

//直接声明和初始化：直接声明map并初始化。
//hash := map[string]int{
//"a": 1,
//"b": 2,
//"c": 3,
//}

//使用 make 函数：使用 make 函数创建一个空的map。
//hash := make(map[string]int)

//先声明后初始化：先声明map，然后使用索引赋值的方式进行初始化。
//var hash map[string]int
//hash = map[string]int{}
//hash["a"] = 1
//hash["b"] = 2

//使用字面值初始化：使用map的零值初始化，然后使用索引赋值的方式进行初始化。
//var hash map[string]int
//hash = map[string]int{}
//hash["a"] = 1
//hash["b"] = 2

//使用映射初始化另一个映射：使用已有的映射初始化另一个映射。
//hash1 := map[string]int{"a": 1, "b": 2}
//hash2 := hash1

func NewHas() map[string]int {

	hash := map[string]int{
		"1": 2,
		"3": 4,
		"5": 6,
	}

	//添加
	hash["9"] = 9

	//删除
	delete(hash, "5")

	//访问
	v, ok := hash["1"]
	fmt.Printf("%d\n", v)
	fmt.Printf("ok = %v\n", ok)
	return hash

}

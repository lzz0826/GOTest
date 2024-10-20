package code

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func JsonT() {
	// 将 Go 结构体转换为 JSON 格式的字节切片
	person := Person{Name: "Alice", Age: 30}
	//物件转JSON
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("JSON marshaling failed:", err)
		return
	}
	fmt.Println("JSON data:", string(jsonData))

	s := string(jsonData)
	// 解析 JSON 字符串为结构体对象
	var person2 Person
	err = json.Unmarshal([]byte(s), &person2)
	if err != nil {
		fmt.Printf(" json.Unmarshal :", err)
	}

	fmt.Println("person2 : ", person2)

	// 将 JSON 格式的字节切片解析为 Go 结构体
	var decodedPerson Person
	err = json.Unmarshal(jsonData, &decodedPerson)
	fmt.Println(reflect.TypeOf(string(jsonData)))
	if err != nil {
		fmt.Println("JSON unmarshaling failed:", err)
		return
	}
	fmt.Println("Decoded person:", decodedPerson)
}

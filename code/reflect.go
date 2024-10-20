package code

import (
	"fmt"
	"reflect"
	"strconv"
)

//反射

//reflect.TypeOf 能获取类型信息；
//reflect.ValueOf 能获取数据的运行时表示；

type Move interface {
	JUP()
	Work(jub string)
}

type Bird struct {
	Name string
	Age  int
}

func (b Bird) JUP() {
	fmt.Println(b.Name + " Jup...")
}

func (b Bird) Work(jub string, time int) {
	fmt.Println(b.Name + " IS Working :" + jub)
	fmt.Println(b.Name + " time is  :" + strconv.Itoa(time))
}

func Reflect() {
	bird := Bird{Name: "Toy", Age: 4}

	birdType := reflect.TypeOf(bird)
	birdValue := reflect.ValueOf(bird)

	fmt.Println(birdType)
	fmt.Println(birdValue)
	// 遍历变量的字段并输出字段名和值

	for i := 0; i < birdType.NumField(); i++ {
		field := birdType.Field(i)
		value := birdValue.Field(i).Interface()
		fmt.Printf("Field %s: %v\n", field.Name, value)
	}

	// 修改变量的值
	// 获取 bird 变量的指针
	birdPtr := &bird
	// 使用反射获取指针指向的值
	birdValue2 := reflect.ValueOf(birdPtr).Elem()

	newValue := "Max"
	newNameField := birdValue2.FieldByName("Name")

	if !newNameField.IsValid() {
		fmt.Println("Field is invalid")
	} else if !newNameField.CanSet() {
		fmt.Println("Field cannot be set")
	}

	if newNameField.IsValid() && newNameField.CanSet() {
		if newNameField.Kind() == reflect.String {
			newNameField.SetString(newValue)
		}
	}

	// 输出修改后的变量
	fmt.Println("Modified dog:", bird)

}

func ReflectMethod() {

	bird := Bird{Name: "Toy", Age: 3}

	// 使用反射获取结构体类型
	dogType := reflect.TypeOf(bird)

	// 遍历结构体的方法
	for i := 0; i < dogType.NumMethod(); i++ {
		method := dogType.Method(i)
		fmt.Println("Method:", method.Name)
	}

	// 使用反射获取结构体值
	birdValue := reflect.ValueOf(bird)

	// 获取指定名称的方法并调用
	barkMethod := birdValue.MethodByName("Work")

	// 准备方法的参数
	args := []reflect.Value{reflect.ValueOf("Ton"), reflect.ValueOf(20)}

	barkMethod.Call(args)
}

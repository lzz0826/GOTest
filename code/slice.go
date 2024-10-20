package code

import (
	_ "fmt"
	_ "strconv"
	_ "time"
)

//// 声明一个空切片
//var dynamicArr []int
//// 向切片中添加元素
//dynamicArr = append(dynamicArr, 1)

//直接声明切片：使用切片字面值直接声明一个切片。
//slice := []int{1, 2, 3}
// 获取整个切片
//result := stack[:]
// 获取从索引 1 开始到最后的所有元素
//result := stack[1:]
// 获取从索引 1 到 3 的所有元素
//result := stack[1:4]
// 获取从索引 0 到 2 的所有元素
//result := stack[:3]
// 获取整个切片的拷贝
//result := append([]string{}, stack...)

//使用 make 函数：使用 make 函数创建一个指定长度和容量的切片。
//slice := make([]int, 3) // 创建一个长度为 3 的切片，初始值为默认值
//slice := make([]int, 3, 5) // 创建一个长度为 3、容量为 5 的切片，初始值为默认值

//从数组或切片中获取子切片：使用切片表达式从数组或切片中获取子切片。
//arr := [5]int{1, 2, 3, 4, 5}
//slice := arr[1:3] // 从数组 arr 中获取索引从 1 到 3（不包含）的子切片

//使用 append 函数：使用 append 函数向切片追加元素，如果切片容量不足会动态扩展切片长度。
//slice := []int{1, 2, 3}
//slice = append(slice, 4) // 追加一个元素到切片末尾

//使用切片初始化另一个切片：通过切片初始化另一个切片，两个切片共享相同的底层数组。
//slice1 := []int{1, 2, 3}
//slice2 := slice1[:] // 初始化一个新切片，与 slice1 共享底层数组

//使用 copy 函数：使用 copy 函数复制一个切片。
//slice1 := []int{1, 2, 3}
//slice2 := make([]int, len(slice1))
//copy(slice2, slice1) // 复制 slice1 到 slice2

func NewSlice() []int {

	//声明
	//slice := []int{1, 2, 3}
	//slice := make([]int, 10)
	//slice[1] = 2

	//创建数组
	arr := [3]int{6, 9, 4}
	//切片 arr[low:high]
	slice := arr[1:3]
	//slice := slice[1:3]
	return slice
}

func NewSliceVstat() []int {
	// 创建一个长度为 3 的整型数组 vstat，并赋值为 {1, 2, 3}
	var vstat [3]int
	vstat[0] = 1
	vstat[1] = 2
	vstat[2] = 3

	// 创建一个指向长度为 3 的整型数组的指针 vauto，并分配内存空间
	var vauto *[3]int = new([3]int)

	// 将 vstat 数组的值复制给 vauto 指针指向的内存空间
	*vauto = vstat

	// 使用切片表达式 vauto[:] 将整型数组转换为切片，并将该切片赋值给 slice 变量
	slice := vauto[:]

	// 返回切片
	return slice
}

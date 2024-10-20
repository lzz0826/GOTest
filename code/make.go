package code

//make 是一个内建函数，用于创建和初始化某些数据结构，
//包括切片 (slice)、映射 (map) 和通道 (channel)。
//make 函数与 new 函数不同，make 不返回指针，而是返回一个初始化后的具体类型。

//切片（Slice）
// 创建一个长度为5，容量为10的整数切片
//numbers := make([]int, 5, 10)

//映射（Map）
// 创建一个初始容量为10的映射，键是整数，值也是整数
//studentScores := make(map[int]int, 10)

//通道（Channel）
// 创建一个缓冲区大小为5的字符串通道
//messages := make(chan string, 5)

//初始长度为 0：当你计划动态地往切片中添加元素时，例如使用 append 函数。
//省略初始长度：当你需要一个已经初始化了指定数量元素的切片，可以直接进行索引访问或替换。
//permits := make([]string, 0, len(permitMap))
//for _, permit := range permitMap {
//    permits = append(permits, permit.Name)
//}

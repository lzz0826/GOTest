package main

import "GOTest/code"

func main() {

	//--------------- := 特例---------------

	code.Notice()

	//--------------- rang for---------------

	//code.RangeArrAndSlice()
	//code.RangeArrMap()
	//code.RangeChar()
	//code.RangeChan()

	//--------------- 切片---------------
	//for i, v := range code.NewSlice() {
	//	fmt.Printf(strconv.Itoa(i))
	//	fmt.Printf(strconv.Itoa(v))
	//}
	//
	//for i, v := range code.NewSliceVstat() {
	//	fmt.Printf(strconv.Itoa(i))
	//	fmt.Printf(strconv.Itoa(v))
	//}

	//----------------hashMap--------------
	//has := code.NewHas()
	//for k, v := range has {
	//	fmt.Printf(k)
	//	fmt.Printf(strconv.Itoa(v))
	//}

	//-----------------JSON转换---------------
	//code.JsonT()

	//----------------interface----------------
	//var rpcErr = interfaceTest.NewError(1, "dd")
	//rpcErr1 := interfaceTest.AsError(rpcErr)
	//fmt.Printf(rpcErr1.Error())

	//grpcError := interfaceTest.GRPCError{Code: 2, Message: "wf"}
	//interfaceTest.PrintInterface(grpcError)

	//实现
	//cat := interfaceTest.Cat{Name: "dvd "}
	//cat.Quack()
	//dog := interfaceTest.Dog{Name: "tt ", Age: 3}
	//dog.Jup()
	//dog.Quack()

	//接口转换成结构体
	//var c interface{} = interfaceTest.Dog{Name: "draven"}
	//switch c.(type) {
	//case interfaceTest.Dog:
	//	cat := c.(interfaceTest.Dog)
	//	cat.Quack()
	//}

	//----------------反射----------------
	//反射属性
	//code.Reflect()
	//反射方法
	//code.ReflectMethod()

	//----------------channel 渠道----------------

	//1. 单向通道（unidirectional channel

	//ch := make(chan string)
	//go func() {
	//	channelTest.UnidirectionalSend(ch, "evew1")
	//	channelTest.UnidirectionalSend(ch, "evew2")
	//	channelTest.UnidirectionalSend(ch, "evew3")
	//	channelTest.UnidirectionalSend(ch, "evew4")
	//	close(ch)
	//}()
	//channelTest.UnidirectionalSendReceive(ch)

	//2. 带缓冲的通道（buffered channel）：
	//channelTest.InitChan()

	//go channelTest.BufferedChannel()
	//channelTest.PChan()
	//channelTest.PChan()
	//time.Sleep(40 * time.Second)

	//3. 使用 select 语句处理多个通道：
	//channelTest.InitCh1Ch2()
	//
	//go channelTest.SelectCh1Ch2()
	//go channelTest.Ch1()
	//go channelTest.Ch2()
	//time.Sleep(4 * time.Second)
	//go channelTest.Ch1()
	////go channelTest.CloseC1()
	//time.Sleep(4 * time.Second)
	//go channelTest.Ch1()
	//time.Sleep(40 * time.Second)

	//4. 使用通道实现同步：
	//channelTest.ChanWorker()
	//time.Sleep(40 * time.Second)

	//Chat 通用 Interface
	//type GenericChannel chan interface{}
	//ch := make(GenericChannel)
	//channelTest.SetChat(ch)
	//channelTest.ChatType(ch)

	//----------------Panic And Recover 异常抛接----------------

	//code.Example(1)

	//fmt.Println("Program continues") // 这行代码会执行，因为在 example 函数中已经进行了 panic 恢复

	//code.DoFinally()
	//----------------Context 上下文----------------

	//code.ContextTest()
	//time.Sleep(40 * time.Second)

	//----------------锁 Lock----------------

	//Mutex 锁 一般用法
	//go syncLock.AddCount(3)
	//go func() {
	//	count := syncLock.GetCount()
	//	fmt.Printf("count : %v", count)
	//}()

	//Mutex 锁 构造
	//mutexStruct := syncLock.MutexStruct{}
	//
	//go mutexStruct.CountAdd()
	//count := mutexStruct.GetCount()
	//fmt.Printf(strconv.Itoa(count))
	//go mutexStruct.CountAdd()
	//count = mutexStruct.GetCount()
	//fmt.Printf(strconv.Itoa(count))
	//go mutexStruct.CountAdd()
	//time.Sleep(4 * time.Second)
	//
	//count = mutexStruct.GetCount()
	//fmt.Printf(strconv.Itoa(count))

	//RmMutex 读写锁
	//mutex := syncLock.RmMutex{Count: 4}
	//
	//for i := 0; i < 5; i++ {
	//	go func() {
	//		for {
	//			mutex.AddCount(1)
	//			time.Sleep(4 * time.Second)
	//		}
	//	}()
	//}
	//
	//for i := 0; i < 5; i++ {
	//	go func() {
	//		for {
	//			count := mutex.GetCount()
	//			fmt.Printf("count: %v", count)
	//			time.Sleep(1 * time.Second)
	//
	//		}
	//	}()
	//}

	//Once 只能执行一次
	//syncLock.OnceTest()
	//for i := 0; i < 5; i++ {
	//	go func() {
	//		syncLock.DoSomeThingR()
	//	}()
	//}
	//syncLock.DoSomeThingR()
	//syncLock.ResetOnce()
	//syncLock.DoSomeThingR()

	//syncLock.DoCond()

	//----------------定时任务 排程----------------
	//cron.CronTest()
	//time.Sleep(40 * time.Second)

}

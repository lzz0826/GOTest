package code

import "fmt"

func Example(cont int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			DoFinally()
		}
	}()

	fmt.Println("Start")

	if cont >= 5 {
		panic("Something went wrong!")
	}
	fmt.Println("End") // 这行代码不会被执行，因为在 panic 之后的代码不会执行

}

// 这里可以模拟 finally
func DoFinally() {
	fmt.Println("DoFinally....")
}

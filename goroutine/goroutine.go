package main

/*
1.不sleep 程序可能直接退出,go程不能执行
2.注意闭包和函数的作用范围和传递方式
*/
import (
	"time"
)

func test() {

	for i := 10; i > 0; i-- {
		go func() {
			println("test ", i)
		}()
	}
	time.Sleep(2 * time.Second)
}
func testVar() {

	for i := 10; i > 0; i-- {
		k := i
		go func() {
			println("testVar ", k)
		}()
	}
	time.Sleep(2 * time.Second)
}
func testFuncVar() {

	for i := 10; i > 0; i-- {
		go func(i int) {
			println("testFuncVar ", i)
		}(i)
	}
	time.Sleep(2 * time.Second)
}

var sleep = make(chan int)

func main() {
	test()
	testVar()
	testFuncVar()
}

/*
连续、交替打印n个A B
A
B
A
B
...
*/

package main

var a = make(chan int)
var b = make(chan int)
var n = 10

func printA() {
	//B等待
	//A执行
	for n > 0 {
		<-b
		println("A")
		a <- 1
	}
}
func printB() {
	//A等待
	//B执行
	for n > 0 {
		<-a
		println("B")
		n -= 1
		b <- 1
	}
}

//启动和等待结束
func whoStart(x string) {
	if x == "A" {
		b <- 1
	} else {
		a <- 1
	}
	for n > 0 {

	}
}
func main() {
	go printA()
	go printB()
	whoStart("B")
}

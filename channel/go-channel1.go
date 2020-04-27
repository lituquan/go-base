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
var n = 1000

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
func main1() {
	go printA()
	go printB()
	whoStart("B")
}

func main() {
	go func() {
		for i := 0; i < n; i++ {
			a <- i
			print("A")
			<-a
		}
	}()
	for i := 0; i < n; i++ {
		<-a
		print("B")
		a <- i
	}
}

package main

var ch = make(chan int)
var n = 10

//生产
func product() {
	for i := 0; i < n; i++ {
		ch <- i
	}
}

//消费
func comsume() int {
	return <-ch
}
func main() {

	go product()

	for i := 0; i < n; i++ {
		println(comsume())
	}
}

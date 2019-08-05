package main

import (
	"sync"
)

var groupStop sync.WaitGroup

func setAll() {
	//设置任务总数
	groupStop.Add(10)
	for i := 0; i < 10; i++ {
		temp := i
		go func() {
			//Done完成一个任务
			defer groupStop.Done()
			println(temp)
		}()
	}
	//等待
	groupStop.Wait()
}

func setOne() {
	for i := 0; i < 10; i++ {
		temp := i
		//设置一个任务
		group.Add(1)
		go func() {
			//Done完成一个任务
			defer group.Done()
			println("setOne", temp)
		}()
	}
	//等待
	group.Wait()
}
func main2() {
	setAll()
	setOne()
}

package main

import (
	"sync"
)

var lock sync.Mutex
var group sync.WaitGroup

func main() {
	mapTemp := make(map[int]int)
	group.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			lock.Lock()
			defer lock.Unlock()
			defer group.Done()
			mapTemp[i] = i
		}(i)
	}

	group.Wait()
	for i := 0; i < 10; i++ {
		func(i int) {
			println(mapTemp[i])
		}(i)
	}
}

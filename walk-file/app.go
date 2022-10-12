package main

import (
	"log"
)

//设置日志格式
func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	//加载文件
	loadCache("./")
	RunServer()
}

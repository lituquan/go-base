package main

//导包
import (
	"encoding/json"    //json解析
	"fmt"              //读写
	"net/http"         //http协议
	_ "net/http/pprof" //服务监控，路由为/debug/pprof，可以观测内存和cpu
)
//打开GC日志   GODEBUG=gctrace=1
//gc 1 @0.002s 12%: 0+2.9+0 ms clock, 0+0/2.9/0+0 ms cpu, 4->7->6 MB, 5 MB goal, 4 P
func main() {
	//返回hello world
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "hello world")
	})
	//返回json
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		word, _ := json.Marshal(map[string]string{"word": "hello world"})
		fmt.Fprint(writer, string(word))
	})
	http.ListenAndServe(":8080", nil)
}

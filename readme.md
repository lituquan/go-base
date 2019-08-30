#### 执行方式 
    打开命令行,cd 进入目录, go run xxx.go
	
#### go常见语法
``````
1.goroutine
	并发操作，只需要发起一个协程
	
2.并发控制
	计数器控制：sync.WaitGroup
	锁控制：sync.Mutex
	通道控制：channel
	
3.协程池化	
	https://blog.csdn.net/Jeanphorn/article/details/79018205	
	worker	工作者
	job 任务	
	dispatcher 分发器，持有jobpool和workerpool
	

4.rpc：
	参考：https://www.jianshu.com/p/74ac2439afb2
	rpc 服务，服务化可以分离不同的业务，但也会带来访问和维护成本。
	
5.http:
	http 发布页面服务或者接口服务，后者可以做前后分离
			
6.命令行调用
	os/exec
	
7.系统信号
	https://studygolang.com/articles/2076
	https://blog.csdn.net/weixin_34015336/article/details/85980854
	https://github.com/tabalt/gracehttp#demo
	
8.优雅重启：socket和process处理
	
9.日志包
	https://www.cnblogs.com/rickiyang/p/11074164.html	



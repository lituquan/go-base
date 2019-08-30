package main

import (
	"fmt"
	"time"
)

//Task
type Task struct{
	f func() error
}
//执行Task
func (t *Task)Execude(){
	t.f()
}
//创建Task
func NewTask(f func() error) Task{
	return Task{
		f:f,
	}
}

//分发器
type Pool struct {
	worker_num int
	JobsChannel chan Task
	stopChannel chan int
}
//构造分发器
func NewPool(cap int,max int)*Pool {
	return &Pool{
		worker_num:cap,
		JobsChannel:make(chan Task,max),
		stopChannel:make(chan int),
	}
}
//建立一个流水线，消化任务
func (p *Pool)worker(work_ID int){
	for task:=range p.JobsChannel{
		task.Execude()
	}
}
//分发器分发打开流水线
func (p *Pool)Run(){
	for i:=0;i<p.worker_num;i++{
		go p.worker(i)//开启流水线
	}
	select {
		case <-p.stopChannel:
			close(p.JobsChannel)
			break
	}
}
//client:task-->Pool.channel
//Pool.worker(s)<--Pool.channel
func rabbit(n int) int{
	if n==1 || n==2{
		return 1
	}else{
		return rabbit(n-1)+rabbit(n-2)
	}
}
func main(){
	target:=20
	pool:=NewPool(4,1000)
	go func() {
		for i:=1;i<=target;i++{
			j:=i
			pool.JobsChannel<-NewTask(func() error {
				fmt.Println(j,rabbit(j))
				return nil
			})
		}
		time.Sleep(2*time.Second)
		pool.stopChannel<-1
	}()
	pool.Run() //消费者
}
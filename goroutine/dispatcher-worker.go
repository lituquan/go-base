package main

import (
	"log"
	"time"
)

//管理者：存储任务
//工作者：执行任务

type Task struct{
	Do func ()
}
//消费者
type Worker struct{
	id int
}

//流水线
func (w *Worker)Run(task Task){
	log.Println(w.id,"正在干活")
	task.Do()
}
//管理者
type Pool struct{
	max int //worker数
	mem chan Task //任务队列
	workers []Worker//消费者们
	stop chan int
}

func (p *Pool)Run(){
	for i:=0;i<p.max;i++{
		p.workers[i]=Worker{i+1}
		w:=p.workers[i]
		go func() {
			for task:=range p.mem{ //分发任务
				w.Run(task)
			}
		}()
	}
	select {
	case <-p.stop: //关闭
		log.Println("stop")
	}
}

func NewPool(maxWorker int,capPool int) *Pool {
	p:=Pool{
		maxWorker,
		make(chan Task,capPool),
		make([]Worker,maxWorker),
		make(chan int),
	}
	return &p
}
func main(){
	p:=NewPool(4,100)
	go func() {
		for i:=100;i<10000;i++{
			temp:=i
			p.mem<-Task{Do: func() { //生产者：提交任务
				sum:=0
				for j:=0;j<=temp;j++{
					sum+=j
				}
				log.Println("计算sum",temp,sum)
			}}
		}
		time.Sleep(5*time.Second)
		p.stop<-1
	}()
	p.Run()
}

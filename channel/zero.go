package main

//https://leetcode-cn.com/problems/print-zero-even-odd/
var a,b,c=make(chan int),make(chan int),make(chan int)

var printNumber=func(a int ){
	print(a)
}
var n=10

func main(){
	go even()
	go odd()
	zero()
		
}

func zero() { 
	for i:=1;i<=n;i++{
		printNumber(0)
		if i&1==0{
			b<-i
		}else{
			c<-i
		}
		<-a
	}
 }  // 仅打印出 0
func even() {
	for ;;{
		select{
			case i:=<-b:
				printNumber(i)
				a<-i
		}
	}
 }  // 仅打印出 偶数
func odd() { 
	for ;;{
		select{
			case i:=<-c:
				printNumber(i)
				a<-i
		}
	}
 }   // 仅打印出 奇数

type ZeroEvenOdd struct{

}

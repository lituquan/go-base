package main

/*
1.exec
	https://blog.csdn.net/youngwhz1/article/details/88662172
2.start,run,wait,output
	https://blog.csdn.net/qq_36874881/article/details/78234005
*/
import (
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

func init()  {
	log.SetFlags(log.Lshortfile)
}
func start(cmdstr string) {
	//命令构造
	cmds := strings.Split(cmdstr, " ")
	cmd := exec.Command(cmds[0], cmds[1:]...)

	stdin,err:=cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	defer stdin.Close()

	//获取输出对象，可以从该对象中读取输出结果
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	defer stdout.Close()
	// 执行命令
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	// 读取输入结果
	_,err=stdin.Write([]byte("hello"))
	if  err != nil {
		log.Fatal(err)
	}else{
		ioutil.ReadAll(stdin)
		log.Println("",)
	}
	// 读取输出结果
	if opBytes, err := ioutil.ReadAll(stdout); err != nil {
		log.Fatal(err)
	} else {
		log.Println(string(opBytes))
	}
}

func run(cmdstr string) {

	//命令构造
	cmds := strings.Split(cmdstr, " ")
	cmd := exec.Command(cmds[0], cmds[1:]...)

	// 执行命令
	if by, err := cmd.Output(); err != nil {
		log.Fatal(err)
	} else {
		log.Println(string(by))
	}

}

func main() {
	start("ping 127.0.0.1 -n 3") //window 使用ping 模拟延迟
	run("ping 127.0.0.1 -n 3")
}

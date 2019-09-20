package main

//将err重定向
import (
	"fmt"
	"os"
	"syscall"
)

var (
	kernel32         = syscall.MustLoadDLL("kernel32.dll")
	procSetStdHandle = kernel32.MustFindProc("SetStdHandle")
)

func SetStdHandle(stdhandle int32, handle syscall.Handle) error {
	r0, _, e1 := syscall.Syscall(procSetStdHandle.Addr(), 2, uintptr(stdhandle), uintptr(handle), 0)
	if r0 == 0 {
		if e1 != 0 {
			return error(e1)
		}
		return syscall.EINVAL
	}
	return nil
}

var f *os.File

func redirect_err() {
	var err error
	f, err = os.Create(`panic.txt`)
	if err != nil {
		fmt.Println("os.Create failed: %v", err)
	}
	err = SetStdHandle(syscall.STD_ERROR_HANDLE, syscall.Handle(f.Fd()))
	if err != nil {
		fmt.Println("SetStdHandle failed: %v", err)
	}
}
func init(){
	redirect_err()
}
func main(){
	fmt.Println("stdout")
	x:=0
	v:=make([]int64,1000/x)
	v=v
	fmt.Println(len(v))

}

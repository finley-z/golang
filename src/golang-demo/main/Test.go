package main

import (
	_"golang-demo/main/socket"
	"golang-demo/main/basic"
	"runtime"
	"fmt"
)

func main()  {
	//socket.Listen()
	basic.TestMutilpArgs(23,"s4t4t4",45.04)
	fmt.Println("cpus:",runtime.NumCPU())
	runtime.GOMAXPROCS(16);

}
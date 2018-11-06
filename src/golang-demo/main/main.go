package main

import (
	_"time"
	_"golang-demo/main/db"
	_"golang-demo/main/web"
	_"fmt"
	_"time"
	_"golang-demo/main/io"
	_"golang-demo/main/basic"
	_"fmt"
	_"golang-demo/main/str"
	_"golang-demo/main/web"
	"golang-demo/main/socket"
	"fmt"
	"io/ioutil"
)

func main() {
	//数据开发测试
	//p := fmt.Println
	//start := time.Now()
	//p(start)
	//
	//
	////db.QueryTest()
	//io.WriteFile()
	//
	//
	//
	//end := time.Now()
	//p(end)
	//web.Index()
	//basic.TestArr()
	//basic.TestSlice()
	//basic.TestMap()
	//var list basic.List
	//list.Add(1)
	//list.Size()
	//list.Remove(1)
	//basic.TestPointer()
	//defer fmt.Println("success")
	//	defer fmt.Println("success")

	//basic.TestError()
	//basic.TestPanic()
	//basic.TestInput()
	//basic.Testvmm()

	//c :=make(chan int)
	//
	//go func() {
	//
	//	fmt.Println("execute routine")
	//	time.Sleep(time.Second*5)
	//	c <- 3
	//}()
	//
	//fmt.Println("c",<-c)
	//str.TestSliceStr()
	//web.Index()
	//socket.Listen()
	con,err :=socket.GetConnet()

	if con.
	result ,err :=ioutil.ReadAll(con)
	if err!=nil{
		fmt.Println("read error",result)

	}
	fmt.Println("buffer",result)

}

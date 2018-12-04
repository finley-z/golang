package main

import (
	_"golang-demo/main/socket"
	_"golang-demo/main/basic"
	_"runtime"
	_"fmt"
	_"golang-demo/main/io"

	_"golang-demo/main/basic"
	_"golang-demo/main/db"
	"golang-demo/main/basic"
)

func main()  {
	//socket.Listen()
	//basic.TestMutilpArgs(23,"s4t4t4",45.04)
	//fmt.Println("cpus:",runtime.NumCPU())
	//runtime.GOMAXPROCS(16);
	//basic.TestArray()

	//basic.TestPrograma()
	//basic.TestMap()
	//basic.TestStruct()
	//basic.TestPanic()
	//basic.TestOs()
	//io.ReadFileInfo()
	//basic.TestFlag()
	//basic.TestStr()
	//basic.TestTime()
	//db.InsertTest()
	//db.TestTrans()

	//user := basic.UserInfo{1, "Allen.Wu", 25}
	//
	//basic.DoFiledAndMethod(user)

	//basic.ParseToJSON(&user)
	//basic.ConvertToObject("{\"Id\":1,\"Name\":\"finley.z\",\"Age\":25}")

	//var val interface{}=23

	//fmt.Println(val.(int))
	basic.AssertType()
}

package basic

import (
	"fmt"
	"strconv"
)


/***
  1、变量、常量声明
  2、流程控制语句（if、for、switch、select）
  3、类型转换
 */


 //常量定义
const defaultPort int=8080
const defaultDes="you have a error message"


//枚举定义
const(
	CREATED=iota
	PAYED
	SUCCESS
	FAIL
)

func TestPrograma()  {

	//测试defer语句
	defer fmt.Println("程序执行完毕")

	fmt.Println("const: the port is",defaultPort)
	fmt.Println("const: the defaultDes is",defaultDes)

	fmt.Println("the enum values is：CREATED=",CREATED,"PAYED=",PAYED,"SUCCESS=",SUCCESS,"FAIL=",FAIL)


	//类型转换
	var f float32=4.1
	var convertf=int(f)
	var v32 int32=324
	//var convertfv32 int64=int64(v32)
	fmt.Println("convertf=",convertf)
	fmt.Println("convertfv32=",int64(v32))


	//数值类型与字符串间转换
	strv:="21434"
	intv,converr:=strconv.Atoi(strv)
	fmt.Println("convert string to int =",intv,converr)


	var sw int=44

	//switch测试
	switch sw {
	case 10:
		fmt.Println("the value ==10")
	case 20:
		fmt.Println("the value ==20")
	case 44:
		fmt.Println("the value ==44")
	default:
		fmt.Println("default value")
	}

	var swstr string="2323"
	switch swstr {
	case "323":
		fmt.Println("the value ==10")
	case "24":
		fmt.Println("the value ==20")
	case "35":
		fmt.Println("the value ==44")
	default:
		fmt.Println("default value")
	}

}

package basic

import (
	"strings"
	"fmt"
)


/**
 字符串测试
 1、字符串常规操作
 2、正则表达式匹配
*/


func TestStr()  {
	//测试子串包含
	str :=strings.Contains("hello world","hello")
	fmt.Println("str contains the sub str",str)

	//测试将字符串转换大小写
	lowLetter :=strings.ToLower("Hello World")
	fmt.Println("lowLetter",lowLetter)

	//测试去除空字符串
	oldStr:="    hello world     "
	fmt.Println("old str:",oldStr,"!")

	trimStr:=strings.TrimSpace("    hello world     ")
	fmt.Println("trimed str:",trimStr,"!")

	fmt.Println("trimed str:","hello","!")
}
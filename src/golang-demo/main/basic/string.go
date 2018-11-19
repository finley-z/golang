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
	str :=strings.Contains("hello world","hello")
	lowLetter :=strings.ToLower("Hello World")
	fmt.Println("str contains the sub str",str)
	fmt.Println("lowLetter",lowLetter)
}
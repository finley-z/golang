package basic

import (
	"strings"
	"fmt"
)

func TestSliceStr()  {
	str :=strings.Contains("hello world","hello")
	lowletter :=strings.ToLower("Hello World")
	fmt.Println("str contains the sub str",str)
	fmt.Println("lowletter",lowletter)
}
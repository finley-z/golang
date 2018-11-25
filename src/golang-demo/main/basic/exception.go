package basic

import (
	"errors"
	"fmt"
)

/**
  panic、recover测试
*/

func TestError()  {
	fmt.Println("teee")
	err :=errors.New("test error")
	fmt.Println(err)
}

func TestPanic()  {
	defer func() {
		//if语句中可以出现初始化语句
		if r:=recover();r!=nil{
			fmt.Println("recover the error:",r)
		}
	}()

	fmt.Println("do recover")

	panic(errors.New("fatal error"))

}
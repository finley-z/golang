package basic

import (
	"errors"
	"fmt"
)

func TestError()  {
	fmt.Println("teee")
	err :=errors.New("test error")
	fmt.Println(err)
}

func TestPanic()  {
	defer func() {
		if r:=recover();r!=nil{
			fmt.Println("recover the error:",r)
		}
	}()

	fmt.Println("do recover")

	panic(errors.New("fatal error"))

}
package basic

import "fmt"


/**
 1、从键盘中输入
 */
func TestInput()  {
	var num string
	fmt.Println("please input number ")


	fmt.Scanln(&num);

	//if err!=nil{
	//	fmt.Println("input error :",err)
	//
	//}
	fmt.Println("input number is:",num)

}
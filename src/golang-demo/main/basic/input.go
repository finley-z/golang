package basic

import "fmt"

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
package basic

import "fmt"

var arrpointer [5]int=[5]int{43,56,65,657,7}

var parr *[5]int=&arrpointer

func TestPointer()  {
	getArraysElemnt(&arrpointer)
}

func getArraysElemnt(arr *[5]int)  {
	var len=len(*arr)
	//var temp=*arr
	for i :=0;i<len;i++ {
		fmt.Println("index:",i,"value:",arr[i])
	}
}

func Testvmm()  {
	var i int =1
	var p *int=&i

	fmt.Println(*p)
}
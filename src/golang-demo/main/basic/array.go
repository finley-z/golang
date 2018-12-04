package basic

import (
	"fmt"
)


/***
数组相关语法与特性测试
*/



//标记下标的数组
var arrayIndexed [1*5]int =[1*5]int{0:2,2:3,1:4,3:5,4:5}

//变长数组  在字面量定义时，确定长度，不可再变
var lengthUnow =[...]rune{'a','汉'}

//多维数组
var duowei [3][4]int=[3][4]int {}

//通过new关键字实例化数组
var newArr *[6]int

func TestArray() {

	for i:=0;i<len(arrayIndexed);i++{
		fmt.Println("index:",i,"value:",arrayIndexed[i])
	}

	//给数组赋值
	duowei[0][2]=666

	for i:=0;i<len(duowei);i++{
		fmt.Println("index i:",i)
		for j:=0;j<len(duowei[i]);j++{
			fmt.Println("    index j:",j,"value:",duowei[i][j])
		}
	}


	newArr=new([6]int)

	newArr[0]=23

	for index,num :=range newArr{
		fmt.Println("index=",index,"num=",num)
	}

}

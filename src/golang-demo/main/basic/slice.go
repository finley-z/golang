package basic

import "fmt"

//字面量创建切片
var slice []int=[]int{2,6,7,7,78,8,8,8,8}

func TestSlice(){

	//通过make关键字创建切片
	var slice2=make([]int,8)
	slice2[0]=3

	//追加内容至切片
	slice2=append(slice2,3,6,7)

	for i:=0;i<len(slice);i++{
		fmt.Println("slice:","index:",i,",",slice[i])
	}

	//从数组中截取切片
	slice1:=slice[:4]

	//获取切片的长度与容量
	fmt.Println("cap slice1:",cap(slice1),"len slice1",len(slice1))

	for j:=0;j<len(slice1);j++{
		fmt.Println("slice1:","index:",j,",",slice1[j])
	}

	fmt.Println(" slice2:",cap(slice2),"len slice1",len(slice2),"vaule%v",slice2)

}
package basic

import "fmt"

var slice []int=[]int{2,6,7,7,78,8,8,8,8}

func TestSlice(){
	for i:=0;i<len(slice);i++{
		fmt.Println("slice:","index:",i,",",slice[i])
	}
	slice1:=slice[:4]

	fmt.Println("cap slice1:",cap(slice1),"len slice1",len(slice1))

	for j:=0;j<len(slice1);j++{
		fmt.Println("slice1:","index:",j,",",slice1[j])
	}
}
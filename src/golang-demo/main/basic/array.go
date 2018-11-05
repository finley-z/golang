package basic

import "fmt"
import "errors"
import "testing"

var arr [1*5]int =[1*5]int{0:2,2:3,1:4,3:5,4:5}
var arqr =[...]int{2,3,4,5,5}
var duowei [3][4]int=[3][4]int {}

var duoweyi =new([3]int)

func TestArr(){
	for i:=0;i<len(arr);i++{
		fmt.Println("index:",i,"value:",arr[i])
	}
	ag :=new([3]int)

	fmt.Println("indexi:",ag)


	for i:=0;i<len(duowei);i++{
		fmt.Println("indexi:",i)
		for j:=0;j<len(duowei[i]);j++{
			fmt.Println("indexj:",j,"value:",duowei[i][j])
		}
		errors.New("ffefe")
	}

}

func testTestArr(t testing.T)  {

}
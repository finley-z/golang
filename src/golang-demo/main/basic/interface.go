package basic

import (
	"fmt"
	_"net/http"
)


/**
 接口与方法
*/


type List []int


type ArrayList interface {
	Size() int
	Add(num int) int
	Remove(index int)int

}

func (list List) Size()  int {
	return len(list)
}

func (list List) Add( num int)  int {
	fmt.Println("add ",num," into array")
	return 1
}

func (list List) Remove( num int)  int {
	fmt.Println("remove ",num," from array")
	return 0
}

func AssertType(){
	var list interface{}=List{1,23}

	m, ok := list.(ArrayList)
	fmt.Println("list type:",m,"yes",ok)
	if _, ok := list.(ArrayList); ok {
		fmt.Printf("value %v of type %T implements ArrayList\n", list, list)
	}

}
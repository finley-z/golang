package basic

import "fmt"

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
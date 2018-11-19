package basic

import (
	"fmt"
)


//创建一个空map
var names map[string] int=map[string] int {}

//通过字面量初始化map
var ids map[int]string =map[int]string {1:"igwgew",2:"r32t",32:"323"}

//通过make关键字创建map
var maps map[string] int=make(map[string]int,40)

func TestMap()  {

	names["finley"]=124
	names["edward"]=124
	names["nick"]=124
	names["john"]=12323

	//删除map指定键
	delete(ids,1)
	ids[1]="tf3r43"

	for i, num := range names {
		fmt.Println("index:", i,"value:",num)
	}

	fmt.Println("map ids vaules=",ids)

	fmt.Println("map maps vaules=",maps)

}
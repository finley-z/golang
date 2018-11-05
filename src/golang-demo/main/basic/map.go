package basic

import (
	"fmt"
)

var names map[string] int=map[string] int {}


var ids map[int]string =map[int]string {1:"igwgew",2:"r32t",32:"323"}
func TestMap()  {
	names["finley"]=124
	names["edward"]=124
	names["nick"]=124
	names["john"]=12323
	delete(names,"john")
	for i, num := range names {
		fmt.Println("index:", i,"value:",num)
	}
}
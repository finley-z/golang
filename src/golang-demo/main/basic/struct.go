package basic

import "fmt"

type User struct {
	id int
	name string
	age int16
	address string
}

func TestStruct()  {
	user :=User{11,"finley",24,"深圳"}
	user1 :=new(User)
	user1.name="334"
	fmt.Println("user :",user)
}
package basic

import "fmt"

func TestMutilpArgs(args ...interface{}) {
	fmt.Println("the first value is ",args[0])
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")

		}
	}
}
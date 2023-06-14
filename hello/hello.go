package main

import "fmt"

func main() {
	str := "Hello"
	fmt.Println(str, "world")

	a := make([]any, 5)
	a[2] = 3
	a[3] = "str"
	fmt.Println(a)

}

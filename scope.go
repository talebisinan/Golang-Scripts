package main

import "fmt"

// x := "foo bar" is not allowed
var x string = "foo bar"

func main() {
	fmt.Println(x)
}

func foo() {
	fmt.Println(x)
}

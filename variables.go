package main

import "fmt"

func main() {
	var x string = "Hey"
	fmt.Println(x)
	x = "Hola"
	fmt.Println(x)
	x = "Merhaba"
	fmt.Println(x)

	var foo string = "foo"
	var bar string = "bar"
	fmt.Println(foo == bar)

	z := 5
	fmt.Println(z)
	var b int
	b = 12
	fmt.Println(b)

	// Not allowed
	// var b || var b;
	// var b int; b = "foo" leads to error

}

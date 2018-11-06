package main

import "fmt"

func main() {
	const foo = "foo"
	fmt.Println(foo)

	// Err: cannot assing to foo
	foo = "bar"
	fmt.Println(foo)
}

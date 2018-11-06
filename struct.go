package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println(p.fname, `says, "Good morning!`)
}

func (s secretAgent) speak() {
	fmt.Println(s.fname, `says, "Shaken, not stirred!`)
}

func main() {
	p1 := person{
		fname: "sinan",
		lname: "talebi",
	}
	fmt.Println(p1)
	p1.speak()

	s1 := secretAgent{
		person{
			"chuck",
			"bartowski",
		},
		true,
	}
	fmt.Println(s1)
	s1.speak()
	s1.person.speak()
}

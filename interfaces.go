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

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person {
		fname: "sinan",
		lname: "talebi",
	}
	s1 := secretAgent {
		person {
			"chuck",
			"bartowski",
		},
		true,
	}
	saySomething(p1)
	saySomething(s1)
}
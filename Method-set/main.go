package main

import (
	"fmt"
)

func main() {

	p1 := person{
		name: "ruben",
	}

	saySomething(&p1)
}

type person struct {
	name string
}

func (p *person) speak() string {
	fmt.Println("la persona:", p.name)
	return p.name
}

type human interface {
	speak() string
}

func saySomething(h human) {
	fmt.Println(h.speak())
}


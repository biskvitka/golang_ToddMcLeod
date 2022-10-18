package main

import (
	"fmt"
)

func main() {
	p := person{"James"}
	p2 := &p
	saySomething(&p)

	saySomething(p2)
}


type person struct {
	first string
}

func (p *person) speak() {
	fmt.Println("hi")
}

type human interface{
	speak()
}

func saySomething(h human) {
	h.speak()
}
package main

import "fmt"

type Speaker interface {
	Speak()
}

type english struct {
}

type korean struct {
}

func (p english) Speak() {
	fmt.Println("welcome")
}

func (p *korean) Speak() {
	fmt.Println("환영")
}

func sayWelcome(s Speaker) {
	s.Speak()
}

func main() {

	e := english{}
	e.Speak()

	k := korean{}
	k.Speak()

	var sk1 Speaker = e
	sk1.Speak()

	var sk2 Speaker = &k
	sk2.Speak()

	sayWelcome(e)
	sayWelcome(&e)

	sayWelcome(&k)

}

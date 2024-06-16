package main

import (
	"fmt"
)

type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

type Lion struct {
	Name string
}

func (D *Dog) Speak() string {
	return (D.Name + " is Dogging")
}

func (L *Lion) Speak() string {
	return (L.Name + " is Lioning")
}

func TellMe(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	d := &Dog{"Rex"}
	l := &Lion{"Leon"}

	TellMe(d)
	TellMe(l)
}

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

func sum(num []int, res chan int) {
	sum := 0

	for _, num := range num {
		sum += num
	}

	res <- sum
}

func main() {
	// d := &Dog{"Rex"}
	// l := &Lion{"Leon"}

	// TellMe(d)
	// TellMe(l)

	numbers := []int{1, 2, 3, 4, 5, 6}

	resultChan := make(chan int)

	mid := len(numbers) / 2
	go sum(numbers[:mid], resultChan)
	go sum(numbers[mid:], resultChan)

	sum1 := <-resultChan
	sum2 := <-resultChan

	fmt.Println(sum1 + sum2)
}

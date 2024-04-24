package main

import "fmt"

type Numbers struct {
	Num1 int
	Num2 int
}

type NumbersInterface interface {
	Sum() int
	Multiply() int
	Substract() int
}

func (n Numbers) Sum() int {
	return n.Num1 + n.Num2
}

func (n Numbers) Multiply() int {
	return n.Num1 * n.Num2
}
func (n Numbers) Substract() int {
	return n.Num1 - n.Num2
}
func main() {
	var i NumbersInterface
	a := Numbers{10, 5}
	i = a

	fmt.Printf(i.Sum())
	fmt.Printf(i.Multiply())
	fmt.Printf(i.Substract())

}

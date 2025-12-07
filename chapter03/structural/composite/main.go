package main

import "fmt"

type Athlete struct {}

func (a *Athlete) Train() {
	fmt.Println("Training")
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim func()
}

func Swim() {
	fmt.Println("Swimming!")
}

type Animal struct {}

func (r *Animal) Eat() {
	println("Eating")
}

type Shark struct {
	Animal
	Swim func()
}

type Swimmer interface {
	Swim()
}

type Trainer interface {
	Train()
}

type SwimmerTmpl struct {}
func (s *SwimmerTmpl) Swim() {
	println("Swimming!")
}

type CompositeSwimmerB struct {
	Trainer
	Swimmer
}

func main() {
	swimmer := CompositeSwimmerB{
		&Athlete{},
		&SwimmerTmpl{},
	}

	swimmer.Train()
	swimmer.Swim()
}
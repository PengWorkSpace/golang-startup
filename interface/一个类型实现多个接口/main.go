package main

import (
	"fmt"
)

type Dog struct {
	name string
}

type Sayer interface {
	say()
}

type Mover interface {
	move()
}

func (dog Dog) move() {
	fmt.Println("i am a dog,i can move")
}

func (dog Dog) say() {
	fmt.Println("i am a cat ,i can say")
}

func main() {

	var sayer Sayer
	var mover Mover
	var dog = Dog{name: "kit"}
	sayer = dog
	mover = dog

	sayer.say()
	mover.move()

}

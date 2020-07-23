package main

import "fmt"

func main() {

	var x Mover
	var wanda = Dog{}
	x = wanda
	wandaShow := x.move()
	fmt.Println(wandaShow)

	//var fugue = &Cat{}
	//x = fugue
	//fugueShow := fugue.move()
	//fmt.Println(fugueShow)

	//无法编译通过
	var peo People = Student{}
	//var peo People = &Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))

}

func (dog Dog) move() string {
	fmt.Println("i am a dog ,i can move")
	return "dog move"
}

func (cat Cat) move() string {
	fmt.Println("i am a cat ,i can move")
	return "cat move"
}

type Mover interface {
	move() string
	eat() string
}

type Cat struct {
}

type Dog struct {
}

func (dog Dog) eat() string {
	panic("implement me")
}

// 面试题
type People interface {
	Speak(string) string
}

type Student struct{}

func (stu Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}
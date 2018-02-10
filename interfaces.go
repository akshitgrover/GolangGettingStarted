package main

import (
	"fmt"
)

type Person interface {
	printJob()
}

type Doctor struct{}

type Teacher struct{}

func main() {

	d := Doctor{}
	t := Teacher{}

	print(d)
	print(t)

}

func (Teacher) printJob() {
	fmt.Println("Teacher")
}

func (Doctor) printJob() {
	fmt.Println("Doctor")
}

func print(i Person) {
	i.printJob()
}

package main

import "time"

func main() {
	go print()
	println("Hello World!")
	time.Sleep(1000)
}

func print() {
	go println("Hello Adder!")
	go println(byte('a'))
}

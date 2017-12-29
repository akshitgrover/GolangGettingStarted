package main

import "fmt"

func main() {
	msg := "Adder"
	sayHello()
	sayHelloByValue(msg)
	sayHelloByReference(&msg)
	println(msg)
	variadic(1, 2, 3, 4, 5, 6, 7)
	variadic()
	fmt.Printf("Sum : %d\n", add(7, 9, 11))
	Val_, Val__ := addRet(7, 9, 11)
	fmt.Printf("Added : %d sum is : %d\n", Val_, Val__)
	Val_, Val__ = addRet_(7, 9, 11)
	fmt.Printf("Added : %d sum is : %d", Val_, Val__)
}

func sayHello() {
	println("Hello Adder!")
}

func sayHelloByValue(msg string) {
	fmt.Printf("Hello %s!\n", msg)
}

func sayHelloByReference(msg *string) {
	fmt.Printf("Hello %s!\n", *msg)
	*msg = "Hello Go!"
}

func variadic(nums ...int) {
	for _, val := range nums {
		println(val)
	}
}

func add(nums ...int) [3]int {
	result := 0
	for _, val := range nums {
		result += val
	}
	arr := [...]int{result, 1, 1}
	return arr
}

func addRet(nums ...int) (int, int) {
	result := 0
	for _, val := range nums {
		result += val
	}
	return len(nums), result
}

func addRet_(nums ...int) (len_ int, sum_ int) {
	for _, val := range nums {
		sum_ += val
	}
	len_ = len(nums)
	return
}

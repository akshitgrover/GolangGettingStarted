package main

import "fmt"

func main() {
	myArray := [...]int{7, 9, 11}
	fmt.Println(myArray)

	myArray_ := [3]int{7, 9, 11}
	fmt.Println(myArray_)

	mySlice := myArray[:]
	mySlice = append(mySlice, 11)
	fmt.Println(mySlice)

	myArray__ := make([]int, 4, 11)
	myArray__ = append(myArray__, 11)
	fmt.Println(myArray__)

	fmt.Println(len(myArray))
	fmt.Println(len(myArray_))
	fmt.Println(len(myArray__))

	myMap := make(map[int]string)
	myMap[7] = "hello adder!"
	myMap[9] = "Hello Adder!"
	myMap[11] = "HELLO ADDER!"
	fmt.Println(myMap[11])
	fmt.Println(myMap)
}

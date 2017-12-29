package main

import "fmt"

func main() {

	myInst := new(myStruct)
	myInst.myField = "myVal"
	fmt.Println(myInst)

	myInst_ := &myStruct{}
	myInst_.myField = "myVal"
	fmt.Println(myInst_)

	myInst__ := newStruct()
	myInst__.myMap[1] = 1
	fmt.Println(myInst__.myMap)

}

type myStruct struct {
	myField string
	myMap   map[int]int
}

func newStruct() *myStruct {

	result := &myStruct{}
	result.myMap = make(map[int]int)

	return result

}

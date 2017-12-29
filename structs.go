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

	myAddInst := newAddStruct([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(myAddInst.findSum(11))
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

type addStruct struct {
	arr [7]int
}

func newAddStruct(arr []int) *addStruct {
	flagStruct := new(addStruct)
	for i := 0; i < 7; i++ {
		flagStruct.arr[i] = arr[i]
	}
	return flagStruct
}

func (asp *addStruct) findSum(offSet int) int {
	result := 0
	for _, val := range asp.arr {
		result += val
	}
	return result + offSet
}

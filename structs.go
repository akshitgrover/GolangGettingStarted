package main

import "fmt"

func main() {

	myInst := new(myStruct)
	myInst.myField = "myVal"
	fmt.Println(myInst)

	myInst_ := myStruct{"myVal"}
	fmt.Println(myInst_)

	myInst__ := &myStruct{}
	myInst__.myField = "myVal"
	fmt.Println(myInst__)

}

type myStruct struct {
	myField string
}

package main

func main() {

	for i := 1; i <= 11; i++ {
		println(i)
	}

	i := 1

	for {
		if i > 11 {
			break
		}
		println(i)
		i++
	}

	var myArray = [3]int{7, 9, 11}

	for idx, val := range myArray {
		println(idx, val)
	}

	var mySlice = make([]int, 11)

	for idx, val := range mySlice {
		println(idx, val)
	}

	var myMap = make(map[int]string)

	myMap[7] = "Seven"
	myMap[9] = "Nine"
	myMap[11] = "Eleven"

	for idx, val := range myMap {
		println(idx, val)
	}

}

package main

func main() {
	var myInt int = 11
	println(myInt)

	myInt_ := 11
	println(myInt_)

	var myFloat32 float32 = 11.
	println(myFloat32)

	var myFloat64 float64 = 11.
	println(myFloat64)

	var myString string = "Hello Adder!"
	println(myString)

	myString_ := "Hello Adder!"
	println(myString_)

	var myComplex = complex(7, 9)
	println(myComplex)
	println(real(myComplex))
	println(imag(myComplex))
}

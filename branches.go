package main

func main() {
	foo := 1

	if foo == 1 {
		println("Hello World!")
	} else {
		println("Hello Adder!")
	}

	if foo = 2; foo == 1 {
		println("Hello World!")
	} else {
		println("Hello Adder!")
	}

	switch foo = 7; foo {
	case 7:
		println("Number Seven!")
	}

	switch foo = 9; {
	case foo > 7:
		println("Number Greater Than Seven!")
	case foo < 7:
		println("Number Less Than Seven!")
	}
}

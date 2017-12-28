package main

const (
	_msg         = iota
	msg   string = "Hello Adder!"
	msg_         = iota
	msg__        = iota
	__msg        = 2 << (1 * iota)
)

func main() {

	const hello = "Hello Adder!"

	println(hello)
	println(msg)
	println(_msg)
	println(msg_)
	println(msg__)
	println(__msg)
}

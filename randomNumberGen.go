package main

import (
	"math/rand"
	"time"
)

func main() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	println(r.Intn(52))

}

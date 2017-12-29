package main

import (
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(7)
	ch := make(chan string)
	chc := make(chan bool)
	go print(ch)
	go printch(ch, chc)
	<-chc
}

func print(ch chan string) {
	for i := byte('a'); i <= byte('z'); i++ {
		ch <- string(i)
	}
	close(ch)
}

func printch(ch chan string, chc chan bool) {
	for i := range ch {
		println(i)
	}
	chc <- true
}

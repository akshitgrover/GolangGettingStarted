package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	scanner.Scan()
	username := scanner.Text()
	fmt.Println(username + " Connected")
	fmt.Fprintln(conn, "Welcome To Go TCP Server, "+username)
	fmt.Println("Connection With " + username + " Closed")
	defer conn.Close()
}

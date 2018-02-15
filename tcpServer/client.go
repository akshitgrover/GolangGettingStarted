package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Enter Your Username: ")
	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString(byte('\n'))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(conn, username)
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		log.Fatalln(err)
	}

	f, _ := os.Create("flag.gotxt")

	io.Copy(f, resp.Body)
}

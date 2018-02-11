package main

import (
	"net/http"
	"os"
)

func main() {
	resp, _ := http.Get("http://google.com")
	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	os.Stdout.Write(bs)
}

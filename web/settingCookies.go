package main

import (
	"net/http"
	"strconv"
)

type visits *http.Cookie

func main() {
	http.HandleFunc("/", Visit)
	http.HandleFunc("/readcount", ReadCount)
	http.ListenAndServe("localhost:8080", nil)
}

func Visit(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		w.Header().Set("Location", "/")
		w.WriteHeader(303)
	}
	d, err := req.Cookie("visits")
	if err != nil {
		var visit visits
		visit = &http.Cookie{}
		visit.Name = "visits"
		visit.Value = "1"
		http.SetCookie(w, visit)
	} else {
		df, _ := strconv.Atoi(d.Value)
		df++
		var visit visits
		visit = &http.Cookie{}
		visit.Name = "visits"
		visit.Value = strconv.Itoa(df)
		http.SetCookie(w, visit)
	}
	message(w)
}

func ReadCount(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		w.Header().Set("Location", "/readcount")
		w.WriteHeader(303)
	}
	d, err := req.Cookie("visits")
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Seomthing Went Wrong"))
	} else {
		w.WriteHeader(200)
		w.Write([]byte("Count Is: " + d.Value))
	}
}

func message(w http.ResponseWriter) {
	w.WriteHeader(200)
	w.Write([]byte("Welcome!"))
}

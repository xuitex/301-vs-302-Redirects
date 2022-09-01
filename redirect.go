package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/permanent", http.RedirectHandler("http://www.google.com", 301))
	http.Handle("/temporary", http.RedirectHandler("https://www.wikipedia.org", 302))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

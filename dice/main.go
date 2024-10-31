package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/roll", roll)

	log.Fatal(http.ListenAndServe(":8088", nil))
}

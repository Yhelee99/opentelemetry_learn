package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func roll(w http.ResponseWriter, r *http.Request) {
	num := 1 + rand.Intn(6)

	_, _ = fmt.Fprintf(w, "reslut: %v", num)
}

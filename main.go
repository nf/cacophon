package main

import (
	"log"
	"net/http"

	_ "github.com/nf/cacophon/backend"
)

func main() {
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// +build !appengine

package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func getCache(req *http.Request) (audio []byte) { return nil }
func putCache(req *http.Request, audio []byte)  {}

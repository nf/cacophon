// +build appengine

package main

import (
	_ "google.golang.org/appengine"
	internal "google.golang.org/appengine/internal"
)

func main() {
	internal.Main()
}

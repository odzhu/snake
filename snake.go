package main

import (
	"net/http"

	snake "github.com/odzhu/snake/snake"
)

var c = new(snake.Whereami)

func main() {
	http.HandleFunc("/", c.Resulthttp)
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"net/http"

	snake "github.com/odzhu/snake/snake"
)

var c = snake.NewCheck()

func handler(w http.ResponseWriter, r *http.Request) {
	c.Name = "testname"
	c.Result = "testResult"
	fmt.Fprintf(w, "App v2 %s!", c)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

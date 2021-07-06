package main

import (
	"fmt"
	"net/http"

	"github.com/tuanti1997qn/go_hello_world/handlers"
)

const Portnumer = ":8080"

func addValue(x, y int) int {
	return x + y
}

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf(fmt.Sprintf("starting application on port %s ", Portnumer))
	_ = http.ListenAndServe(Portnumer, nil)
}

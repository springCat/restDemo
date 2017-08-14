package main

import (
	"fmt"
	"net/http"
		"./hello"
)

func main() {
	r := hello.NewHelloResource()

	http.HandleFunc("/", r.Handler)
	err := http.ListenAndServe(":3000", nil)
	fmt.Println(err)
}

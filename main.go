package main

import (
	"fmt"
	"net/http"
)

func main() {
	r := NewHelloResource()

	http.HandleFunc("/", r.Handler)
	err := http.ListenAndServe(":3000", nil)
	fmt.Println(err)
}

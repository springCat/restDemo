package main

import "net/http"
import (
	r "./resource"
	"fmt"
)

type HelloResource struct {
	*r.RestResource
}

var channel = make(chan int, 1)
var count = 0

func NewHelloResource() r.HttpResource {
	resource := r.New()
	resource.Get(func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte("i am get respose"))
	})

	resource.Post(func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte("i am post respose"))
	})

	resource.Put(func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte("i am put respose"))
	})

	resource.Delete(func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte("i am delete respose"))
	})
	return resource
}

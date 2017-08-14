package resource

import "net/http"

type Resource interface {
	Get(f func(resp http.ResponseWriter, req *http.Request))
	Post(f func(resp http.ResponseWriter, req *http.Request))
	Put(f func(resp http.ResponseWriter, req *http.Request))
	Delete(f func(resp http.ResponseWriter, req *http.Request))
	Handler(resp http.ResponseWriter, req *http.Request)
}

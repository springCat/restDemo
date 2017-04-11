package resource

import "net/http"

type RestResource struct {
	methods map[string]func(http.ResponseWriter, *http.Request)
}

func (resource RestResource) init() {
	resource.methods = make(map[string]func(resp http.ResponseWriter, req *http.Request), 10)
}

func (resource RestResource) Get(f func(resp http.ResponseWriter, req *http.Request)) {
	resource.methods["GET"] = f
}

func (resource RestResource) Post(f func(resp http.ResponseWriter, req *http.Request)) {
	resource.methods["POST"] = f
}

func (resource RestResource) Put(f func(resp http.ResponseWriter, req *http.Request)) {
	resource.methods["PUT"] = f
}

func (resource RestResource) Delete(f func(resp http.ResponseWriter, req *http.Request)) {
	resource.methods["DELETE"] = f
}

func (resource RestResource) Handler(resp http.ResponseWriter, req *http.Request) {
	methodStr := req.Method
	if methodStr == "" {
		resp.WriteHeader(403)
		resp.Write([]byte("http method error"))
		return
	}
	method := resource.methods[methodStr]
	if method == nil {
		resp.WriteHeader(404)
		resp.Write([]byte("no this method"))
		return
	}
	method(resp, req)
}

func New() *RestResource {
	return &RestResource{
		methods: make(map[string]func(resp http.ResponseWriter, req *http.Request), 10),
	}
}

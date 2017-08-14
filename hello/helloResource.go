package hello

import "net/http"
import (
		r "../resource"
		"encoding/json"
		"encoding/xml"
		"html/template"
)

type HelloResource struct {
		*r.RestResource
}

type Result struct {
		Code int
		Msg  string
}

func NewHelloResource() r.Resource {
		resource := *(r.NewRestResource())

		resource.Get(func(resp http.ResponseWriter, req *http.Request) {

				result := &Result{
						Code: 200,
						Msg:  "i am get respose",
				}

				body := `
				<!DOCTYPE html>
				<html>
				<head>
				<meta charset="UTF-8">
				<title>这是一个HTML网页</title>
				</head>
				<body>
				<p>Hello HTML5</p>
				<p>{{.Code}}</p>
				<p>{{.Msg}}</p>
				</body>
				</html>
				`
				t := template.New("html example")
				html, err := t.Parse(body)
				if err == nil {
						html.Execute(resp, result)
				}
		})

		resource.Post(func(resp http.ResponseWriter, req *http.Request) {
				result := &Result{
						Code: 20000,
						Msg:  "not support post method",
				}
				v, error := json.Marshal(result)
				if error == nil {
						resp.WriteHeader(500)
						resp.Write(v)
				}
		})

		resource.Put(func(resp http.ResponseWriter, req *http.Request) {
				result := &Result{
						Code: 200,
						Msg:  "i am put respose",
				}
				v, error := xml.Marshal(result)
				if error == nil {
						resp.Write(v)
				}
		})

		resource.Delete(func(resp http.ResponseWriter, req *http.Request) {
				result := &Result{
						Code: 200,
						Msg:  "i am delete respose",
				}
				v, error := json.Marshal(result)
				if error == nil {
						resp.Write(v)
				}
		})
		return resource
}

package main

import (
	"github.com/azer/atlas"
	"time"
)

var api = atlas.New(&atlas.URLs{
	"/user/:name/:surname": User,
	"/company/:id":         Company,
	"/hello":               Hello,
	"/now":                 Now,
	"/querystring":         QueryString,
	"/post":                PostName,
	"/json-post":           JSONPost,
	"/ip":                  IP,
	"/html":                HTML,
})

type Person struct{ Name, Surname string }

func main() {
	api.Start(":8080")
}

func Hello(request *atlas.Request) *atlas.Response {
	return atlas.Success("Hello World")
}

func Now(request *atlas.Request) *atlas.Response {
	return atlas.Success(time.Now())
}

func User(request *atlas.Request) *atlas.Response {
	return atlas.Success(&Person{request.Params["name"], request.Params["surname"]})
}

func Company(request *atlas.Request) *atlas.Response {
	return atlas.Error(500, "You should see an error.")
}

func QueryString(request *atlas.Request) *atlas.Response {
	return atlas.Success(request.Query)
}

func PostName(request *atlas.Request) *atlas.Response {
	return atlas.Success(request.Form.Get("name"))
}

func JSONPost(request *atlas.Request) *atlas.Response {
	var data map[string]string
	err := request.JSONPost(&data)

	if err != nil {
		return atlas.Error(500, err)
	}

	return atlas.Success(data)
}

func IP(request *atlas.Request) *atlas.Response {
	return atlas.Success(request.RemoteAddr)
}

func HTML(request *atlas.Request) *atlas.Response {
	return atlas.Custom([]byte("<h1>hello world</h1>"))
}

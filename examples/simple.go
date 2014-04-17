package main

import (
	"github.com/azer/atlas"
	"time"
)

var api = atlas.New(atlas.Map{
	"/user/:name/:surname": User,
	"/company/:id":         Company,
	"/hello":               Hello,
	"/now":                 Now,
	"/querystring":         QueryString,
	"/post":                PostName,
	"/json-post":           JSONPost,
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
	return atlas.Error(500, "An error occured")
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

	return atlas.Success(data["name"])
}

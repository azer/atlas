## Atlas

Minimalistic Go Library for Creating JSON API Servers.

```go
import "github.com/azer/atlas"

var api = atlas.New(atlas.Map{ "/": Hello })

func Hello(request *atlas.Request) *atlas.Response {
	return atlas.Success("Hello World")
}

api.Start(":8080")
```

It'll output the JSON encoding of whatever is returned from request handlers:

```bash
$ curl localhost:8080
{ ok: true, result: "Hello World" }
```

JSONP is enabled by default:

```bash
$ curl localhost:8080?callback=foobar
foobar({ ok: true, result: "Hello World" })
```

## Install

```bash
$ go get github.com/azer/atlas
```

## Manual

Create a new API by defining Sinatra-like routes:

```go
import "github.com/azer/atlas"

var api = atlas.New(atlas.Map{
	"/user/:name/:surname": User,
	"/company/:id": Company,
	"/hello": Hello,
})
```

Every route points to a function (atlas.Handler) that takes [*atlas.Request](#requests) and returns a `*Response`. Atlas comes with `Success` and `Error` functions that converts anything to a response:

```go
func Hello(request *atlas.Request) *atlas.Response {
	return atlas.Success("Hello World")
}
```

You can start the API by calling `Start` method:

```go
api.Start(":8080")
```

Atlas will output JSON for you:

```bash
$ curl http://localhost:8080/hello
{ ok: true, result: "Hello World!" }
```

If you leave "/" without a handler, Atlas will show a simple API index by default:

```bash
$ curl http://localhost:8080/
{
  "welcome": true,
  "endpoints": [
    "/company/:id",
    "/hello",
    "/user/:name/:surname"
  ]
}
```

You can return structs, maps, etc... Atlas will output them as JSON for you:

```go
type Person struct { Name, Surname string }

func User(request *atlas.Request) *atlas.Response {
	return atlas.Success(&Person{request.Params["name"], request.Params["surname"]})
}
```

Requests to `/user/:name/:surname` will get:

```bash
$ curl http://localhost:8080/user/john/smith
{
  ok: true,
  result: {
    Name: "john",
    Surname: "smith"
  }
}
```

Atlas also simplifies JSON form posts:

```go
func HelloWorld (request *atlas.Request) *atlas.Response {
  var data map[string]string
  err := request.JSONPost(&data)
  
  if err != nil {
    return atlas.Error(500, err)
  }
  
  return atlas.Success(data)
}
```

### Using JSON Tags

If you'd like to modify struct keys for API, here is an example of how to do it:

```go
type Person struct {
  Name string `json:"name"`
  Surname string `json:"surname"`
}
```

### Error Handling

Atlas also has a handy method to produce errors:

```go
func Company(request *atlas.Request) *atlas.Response {
	return atlas.Error(500, "Not Implemented Yet")
}
```

Will output:

```bash
$ curl localhost:8080/company/foobar
{
  error: "Not implemented yet"
}
```

Checkout `examples/` for more info.

## Reference

### atlas.Request

Atlas will pass you its modified version of requests:

```go
type Request struct {
	Header map[string][]string
	Params urlrouter.Params

	Method string
	Host   string
	URL    *url.URL
	GET    bool
	POST   bool

	Form     url.Values
	PostForm url.Values
	Query    url.Values
}
```

### atlas.Manual

If you'd like to structure the entire response, including the wrapper, use `Manual` instead of `Success` or `Error`.
Here is an example of how to do that:

```go
func Hello(request *atlas.Request) *atlas.Response {
	return atlas.Manual(200, "Hello World")
}
```

### atlas.API.Listen

To provide your own net.Listener, call `Listen`:

```go
listener, _ := net.Listen("tcp", "0.0.0.0:6666")

var api = atlas.New(atlas.Map{ "/": Hello })
api.Listen(listener)
```

## Debugging

Atlas uses [debug](http://github.com/azer/debug) for logging. Enable verbose mode by:

```bash
DEBUG=* go run my server
```

![](https://i.cloudup.com/8uRNKNk9I2.png)

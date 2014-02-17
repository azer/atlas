## Atlas

Minimalistic Go Library for Creating JSON API Servers.

```go
import "github.com/azer/atlas"

var api = atlas.New(atlas.Map{ "/": Hello })

func Hello(request *atlas.Request) *atlas.Response {
	return atlas.Success("Hello World")
}

api.start(":8080")
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

Or you can provide your own net.Listener and call `Listen`:

```go
listener, _ := net.Listen("tcp", "0.0.0.0:6666")
api.Listen(listener)
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
	return atlas.Error(500, "An error occured")
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

### Requests

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
}
```

## Debugging

Atlas uses [debug](http://github.com/azer/debug) for logging. Enable verbose mode by:

```bash
DEBUG=* go run my server
```

![](https://i.cloudup.com/8uRNKNk9I2.png)

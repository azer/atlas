## Atlas

Minimalistic JSON API Server for Go.

## Install

```bash
$ go get github.com/azer/atlas
```

## Usage

Create a new API by defining Sinatra-like routes:

```go
import (
  "github.com/azer/atlas"
  "errors"
)

var api = atlas.API(atlas.Map{
        "/person/:name/:surname": Person,
        "/company/:id": Company,
        "/hello": Hello,
})
```

Every route points to a function (atlas.Handler):

```go
func Home (request *atlas.Request) *atlas.Response {
  return atlas.Success("Hello World!")
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

func User (request *atlas.Request) *atlas.Response {
  return atlas.Success(&Person{request.params.name, request.params.surname})
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

```
func Company (request *atlas.Request) *atlas.Response {
  return atlas.Error(200, "Not implemented yet")
}
```

Will output:

```bash
$ curl localhost:8080/company/foobar
{
  error: "Not implemented yet"
}
```

## Debugging

Atlas uses [debug](http://github.com/azer/debug) for logging. Enable verbose mode by:

```bash
DEBUG=* go run my server
```

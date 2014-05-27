package atlas

import (
	"github.com/azer/url-router"
	"net/http"
	"net/url"
	"encoding/json"
)

type Request struct {
	*http.Request
	Params urlrouter.Params

	GET    bool
	POST   bool

	Query    url.Values
}

func (request *Request) JSONPost(value interface{}) error {
	return json.NewDecoder(request.Body).Decode(&value)
}

func NewRequest(request *http.Request, params urlrouter.Params) *Request {
	query, _ := url.ParseQuery(request.URL.RawQuery)

	request.ParseForm()
	return &Request{
		Request: request,
		Params: params,
		GET: request.Method == "GET",
		POST: request.Method == "POST",
		Query: query,
	}
}

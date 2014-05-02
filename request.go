package atlas

import (
	"github.com/azer/url-router"
	"net/http"
	"net/url"
	"encoding/json"
	"io"
)

type Request struct {
	Header map[string][]string
	Params urlrouter.Params

	Method string
	Host   string
	URL    *url.URL
	GET    bool
	POST   bool

	Body     io.ReadCloser
	Form     url.Values
	PostForm url.Values
	Query    url.Values
}

func (request *Request) JSONPost(value interface{}) error {
	return json.NewDecoder(request.Body).Decode(&value)
}

func NewRequest(request *http.Request, params urlrouter.Params) *Request {
	query, _ := url.ParseQuery(request.URL.RawQuery)

	request.ParseForm()
	return &Request{
		request.Header,
		params,

		request.Method,
		request.Host,
		request.URL,
		request.Method == "GET",
		request.Method == "POST",

		request.Body,
		request.Form,
		request.PostForm,
		query,
	}
}

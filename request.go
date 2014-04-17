package atlas

import (
	"github.com/azer/url-router"
	"net/http"
	"net/url"
	"encoding/json"
)

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

func (request *Request) JSONPost(value interface{}) error {
	for key, _ := range request.Form {
		err := json.Unmarshal([]byte(key), value)

		if err != nil {
			return err
		}

		break
	}

	return nil
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

		request.Form,
		request.PostForm,
		query,
	}
}

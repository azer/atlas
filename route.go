package atlas

import (
	"github.com/azer/url-router"
	"net/http"
)

func (api *API) Route(w http.ResponseWriter, r *http.Request) {
	var params urlrouter.Params

	match := api.Router.Match(r.URL.Path)

	if match == nil {
		params = urlrouter.Params{}
	} else {
		params = match.Params
	}

	request := NewRequest(r, params)

	if match == nil && r.URL.Path == "/" {
		api.Print(w, request, api.Index)
		return
	}

	if match == nil {
		log.Error("No routes match %s", r.URL.Path)
		api.Print(w, request, NotFound)
		return
	}

	handler := (*api.URLs)[match.Pattern]

	api.Print(w, request, handler(request))
}

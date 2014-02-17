package atlas

import "net/http"

func (api *API) Route(w http.ResponseWriter, r *http.Request) {
	match := api.Router.Match(r.URL.Path)

	if match == nil && r.URL.Path == "/" {
		api.Print(w, api.Index)
		return
	}

	if match == nil {
		debug("Unable to match %s", r.URL.Path)
		api.Print(w, NotFound)
		return
	}

	debug("Matched %s with %s", r.URL.Path, match.Pattern)

	request := NewRequest(r, match.Params)
	handler := api.URLs[match.Pattern]

	api.Print(w, handler(request))
}

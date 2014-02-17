package atlas

import "github.com/azer/url-router"

func New (urls Map) *API {
	debug("Initializing a new API server...")

	server := &Server{}
	index := NewIndex(urls)
	router := urlrouter.New()

	for pattern, _ := range urls {
		router.Add(pattern)
	}

	api := &API{
		urls,
		*router,
		server,
		index,
	}

	server.Route = api.Route

	return api
}

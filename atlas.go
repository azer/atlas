package atlas

import (
	"github.com/azer/logger"
	"github.com/azer/url-router"
)

var log = logger.New("atlas")

func New(urls Map) *API {
	log.Info("Initializing a new API server...")

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

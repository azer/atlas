package atlas

import "net/http"

type Map map[string]Handler
type Handler func(request *Request)*Response

type Server struct {
	Route func(http.ResponseWriter, *http.Request)
}

func (s *Server) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	s.Route(w, r)
}

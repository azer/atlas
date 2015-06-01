package atlas

import (
	"fmt"
	"github.com/azer/url-router"
	"net"
	"net/http"
)

type API struct {
	URLs   Map
	Router urlrouter.Router
	Server *Server
	Index  *Response
}

func (api *API) Start(addr string) {
	log.Info("Starting on %s", addr)
	http.ListenAndServe(addr, api.Server)
}

func (api *API) Listen(listener net.Listener) {
	log.Info("Listening on %s", listener.Addr())
	http.Serve(listener, api.Server)
}

func (api *API) Print(writer http.ResponseWriter, request *Request, response *Response) {
	writer.WriteHeader(response.Code)
	output := response.Stringify()

	if callback, jsonp := request.Query["callback"]; jsonp {
		output = fmt.Sprintf("%s(%s)", callback[0], output)
	}

	fmt.Fprintf(writer, output)
}

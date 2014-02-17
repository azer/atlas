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
	debug("Starting on %s", addr)
	http.ListenAndServe(addr, api.Server)
}

func (api *API) Listen(listener net.Listener) {
	debug("Listening on %s", listener.Addr())
	http.Serve(listener, api.Server)
}

func (api *API) Print(writer http.ResponseWriter, response *Response) {
	writer.WriteHeader(response.Code)
	fmt.Fprintf(writer, response.Stringify())
}

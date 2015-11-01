package atlas

import (
	"fmt"
	"github.com/azer/url-router"
	"net"
	"net/http"
)

type API struct {
	URLs   *URLs
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
	if response.IsJSON {
		api.PrintJSON(writer, request, response)
		return
	}

	contentType := http.DetectContentType(response.Output)

	writer.WriteHeader(response.Code)
	writer.Header().Set("Content-Type", contentType)
	writer.Write(response.Output)
}

func (api *API) PrintJSON(writer http.ResponseWriter, request *Request, response *Response) {
	output, err := response.JSON()
	contentType := "application/json"

	if err != nil {
		output, _ = InternalError.JSON()
		return
	}

	if callback, isJSONP := request.Query["callback"]; isJSONP {
		contentType = "application/javascript"
		output = fmt.Sprintf("%s(%s)", callback[0], output)
	}

	writer.WriteHeader(response.Code)
	writer.Header().Set("Content-Type", contentType)
	fmt.Fprintf(writer, output)
}

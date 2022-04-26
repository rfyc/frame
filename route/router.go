package route

import (
	"net/http"
)

type Router struct {
	URI  URIHandler
	HTTP HTTPHandler
}

func (this *Router) ServeHTTP(response http.ResponseWriter, request *http.Request) {

	//******** found exec ********//
	execController, execAction := this.URI.Parse(request.URL.Path)
	if execController == nil || execAction == nil {
		response.Write([]byte(request.URL.Path + " not found"))
		return
	}
	handler := this.HTTP.New(request, response)
	*execController.In() = *handler.In()
	execController.Ctx(request.Context())
	execAction.Ctx(request.Context())
	execController.Run(execAction)
	handler.Out(execController.Out())
}

func (this *Router) ServeTCP() {

}

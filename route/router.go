package route

import (
	"net/http"
)

type Router struct {
	Regsiter RegistHandler
	HTTP     HTTPHandler
}

func (this *Router) ServeHTTP(response http.ResponseWriter, request *http.Request) {

	//find exec
	execController, execAction := this.Regsiter.Parse(request.URL.Path)
	if execController == nil || execAction == nil {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte("not_found"))
		return
	}

	//nit exec
	execAction.Init(request.Context())
	handler := this.HTTP.New(request, response)
	execController.Init(request.Context(), handler.In(), execAction)

	//check exec
	if err := execController.Prepare(); err != nil {
		handler.Out(execController.Out(err, nil))
		return
	}

	//check exec
	if err := execAction.Prepare(); err != nil {
		handler.Out(execController.Out(err, nil))
		return
	}

	//run exec
	handler.Out(execController.Out(execAction.Run()))
}

func (this *Router) ServeTCP() {

}

package route

import (
	"net/http"
)

type Handler struct {
	IRoute
	IHttpHandler
}

func (this *Handler) ServeHTTP(response http.ResponseWriter, request *http.Request) {

	//******** found exec ********//
	execController, execAction := this.Parse(request.URL.Path)
	if execController == nil || execAction == nil {
		response.Write([]byte(request.URL.Path + " not found"))
		return
	}
	this.IHttpHandler.Init(request, response)
	*execController.Input() = *this.IHttpHandler.Input()
	execController.Ctx(request.Context())
	execAction.Ctx(request.Context())
	code, content := execController.RunAction(execAction.Run)
	response.WriteHeader(code)
	response.Header()
	response.Write(content)
}

func (this *Handler) ServeTCP() {

}

func ServeHTTP() {

}

package route

import (
	"net/http"
	"time"
)

type Handler struct {
	router
	*handlerTCP
	*handlerHTTP
}

func (this *Handler) ServeHTTP(response http.ResponseWriter, request *http.Request) {

	//******** found exec ********//
	execController, execAction := this.parse(request.URL.Path)
	if execController == nil || execAction == nil {
		response.Write([]byte(request.URL.Path + " not found"))
		return
	}
	this.handlerHTTP = &handlerHTTP{
		request:        request,
		response:       response,
		btime:          time.Now(),
		execController: execController,
		execAction:     execAction,
	}
}

func (this *Handler) ServeTCP() {

}

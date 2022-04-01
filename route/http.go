package route

import (
	"net/http"
	"time"
)

type IHttpHandler interface {
	Init(*http.Request, http.ResponseWriter)
	Input() *Input
	Output(*Output)
}

type handlerHTTP struct {
	request  *http.Request
	response http.ResponseWriter
	begin    time.Time
}

func (this *handlerHTTP) Init(request *http.Request, response http.ResponseWriter) {
	this.request = request
	this.response = response
	this.begin = time.Now()
}

func (this *handlerHTTP) Input() *Input {
	input := &Input{}
	return input
}

func (this *handlerHTTP) Output(out *Output) {
	this.response.Header()
	this.response.Write(out.Content)
}

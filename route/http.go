package route

import (
	"net/http"
	"time"
)

type HTTPHandler interface {
	New(*http.Request, http.ResponseWriter) HTTPHandler
	In() *Input
	Out(*Output)
}

type DefaultHTTP struct {
	request  *http.Request
	response http.ResponseWriter
	begin    time.Time
}

func (this *DefaultHTTP) New(request *http.Request, response http.ResponseWriter) HTTPHandler {
	return &DefaultHTTP{
		request:  request,
		response: response,
		begin:    time.Now(),
	}
}

func (this *DefaultHTTP) In() *Input {
	input := &Input{}

	//******** request get ********//
	for key, value := range this.request.URL.Query() {
		count := len(value)
		if count > 1 {
			input.Request[key] = value
			input.Get[key] = value
		} else if count == 1 {
			input.Request[key] = value[0]
			input.Get[key] = value[0]
		}
	}

	//******** request post ********//
	this.request.ParseForm()
	for key, value := range this.request.PostForm {
		count := len(value)
		if count > 1 {
			input.Request[key] = value
			input.Post[key] = value
		} else if count == 1 {
			input.Request[key] = value[0]
			input.Post[key] = value[0]
		}
	}

	//******** header ********//
	for key, value := range this.request.Header {
		count := len(value)
		if count > 1 {
			input.Header[key] = value
		} else if count == 1 {
			input.Header[key] = value[0]
		}
	}

	//******** cookie ********//
	for _, cookie := range this.request.Cookies() {
		input.Cookie[cookie.Name] = cookie.Value
	}

	//******** server ********//
	input.Server.IsHTTP = true
	input.Server.IsPOST = this.request.Method == "POST"
	input.Server.IsGET = this.request.Method == "GET"
	input.Server.IsAJAX = len(this.request.Header.Get("x-requested-with")) > 0
	input.Server.RemoteAddr = this.request.RemoteAddr
	input.Server.HostName = this.request.URL.Hostname()
	input.Server.Port = this.request.URL.Port()
	input.Server.Path = this.request.URL.Path
	input.Server.Query = this.request.URL.RawQuery
	input.Server.Referer = this.request.Referer()
	input.Server.UserAgent = this.request.UserAgent()

	return input
}

func (this *DefaultHTTP) Out(out *Output) {
	this.response.Header()
	this.response.Write(out.Content)
}

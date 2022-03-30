package ctx

import (
	"context"
	"net/http"
)

func New() *Ctx {
	return &Ctx{
		Input: &Input{
			Request: map[string]interface{}{},
			Get:     map[string]interface{}{},
			Post:    map[string]interface{}{},
			Cookie:  map[string]interface{}{},
			Header:  map[string]interface{}{},
		},
		Output: &Output{
			Headers: map[string]string{},
		},
		Server: &Server{},
	}
}

type Ctx struct {
	Action     string
	Controller string
	Context    context.Context
	Input      *Input
	Output     *Output
	Server     *Server
}

func (this *Ctx) SetHTTP(request *http.Request) {

	//******** context ********//
	this.Context = request.Context()

	//******** request get ********//
	for key, value := range request.URL.Query() {
		count := len(value)
		if count > 1 {
			this.Input.Request[key] = value
			this.Input.Get[key] = value
		} else if count == 1 {
			this.Input.Request[key] = value[0]
			this.Input.Get[key] = value[0]
		}
	}

	//******** request post ********//
	request.ParseForm()
	for key, value := range request.PostForm {
		count := len(value)
		if count > 1 {
			this.Input.Request[key] = value
			this.Input.Post[key] = value
		} else if count == 1 {
			this.Input.Request[key] = value[0]
			this.Input.Post[key] = value[0]
		}
	}

	//******** header ********//
	for key, value := range request.Header {
		count := len(value)
		if count > 1 {
			this.Input.Header[key] = value
		} else if count == 1 {
			this.Input.Header[key] = value[0]
		}
	}

	//******** cookie ********//
	for _, cookie := range request.Cookies() {
		this.Input.Cookie[cookie.Name] = cookie.Value
	}

	//******** server ********//
	this.Server.IsHTTP = true
	this.Server.IsPOST = request.Method == "POST"
	this.Server.IsGET = request.Method == "GET"
	this.Server.IsAJAX = len(request.Header.Get("x-requested-with")) > 0
	this.Server.RemoteAddr = request.RemoteAddr
	this.Server.HostName = request.URL.Hostname()
	this.Server.Port = request.URL.Port()
	this.Server.Path = request.URL.Path
	this.Server.Query = request.URL.RawQuery
	this.Server.Referer = request.Referer()
	this.Server.UserAgent = request.UserAgent()
}

func (this *Ctx) SetTCP() {

}

type Input struct {
	Request map[string]interface{}
	Get     map[string]interface{}
	Post    map[string]interface{}
	Cookie  map[string]interface{}
	Header  map[string]interface{}
}

type Output struct {
	Status  string
	Error   string
	Content []byte
	Cookies []*http.Cookie
	Headers map[string]string
}

func (this *Output) Redirect(uri string) {
	this.Headers["Location"] = uri
}

type Server struct {
	IsGET      bool
	IsPOST     bool
	IsAJAX     bool
	IsHTTP     bool
	IsHTTPS    bool
	IsTCP      bool
	IsUDP      bool
	HostName   string
	Port       string
	Path       string
	Query      string
	Referer    string
	UserAgent  string
	RemoteAddr string
}

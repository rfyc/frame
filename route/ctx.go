package route

import (
	"net/http"
)

func SetHTTP(input *Input, request *http.Request) {

	//******** request get ********//
	for key, value := range request.URL.Query() {
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
	request.ParseForm()
	for key, value := range request.PostForm {
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
	for key, value := range request.Header {
		count := len(value)
		if count > 1 {
			input.Header[key] = value
		} else if count == 1 {
			input.Header[key] = value[0]
		}
	}
	//******** cookie ********//
	for _, cookie := range request.Cookies() {
		input.Cookie[cookie.Name] = cookie.Value
	}
	//******** server ********//
	input.Server.IsHTTP = true
	input.Server.IsPOST = request.Method == "POST"
	input.Server.IsGET = request.Method == "GET"
	input.Server.IsAJAX = len(request.Header.Get("x-requested-with")) > 0
	input.Server.RemoteAddr = request.RemoteAddr
	input.Server.HostName = request.URL.Hostname()
	input.Server.Port = request.URL.Port()
	input.Server.Path = request.URL.Path
	input.Server.Query = request.URL.RawQuery
	input.Server.Referer = request.Referer()
	input.Server.UserAgent = request.UserAgent()
}

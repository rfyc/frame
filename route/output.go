package route

import "net/http"

type Output struct {
	Code    string
	Error   error
	Content []byte
	Cookies []*http.Cookie
	Headers map[string]string
}

func (this *Output) Redirect(uri string) {
	this.Headers["Location"] = uri
}

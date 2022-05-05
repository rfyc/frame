package connect

import (
	"errors"
	"github.com/rfyc/frame/utils/conv"
	"github.com/rfyc/frame/utils/file"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

type HTTPServer struct {
	Address  string
	Static   string
	GraceFul bool
	listener net.Listener
	server   *http.Server
	serveMux *http.ServeMux
}

func (this *HTTPServer) Init() {

	this.serveMux = http.NewServeMux()
	this.server = &http.Server{Handler: this.serveMux}
}

func (this *HTTPServer) listen(addr string) (net.Listener, error) {

	fd_ptr := conv.Int(os.Getenv(addr))

	if this.GraceFul && fd_ptr > 0 {
		fd := os.NewFile(uintptr(fd_ptr), "")
		return net.FileListener(fd)
	}

	return net.Listen("tcp", addr)
}

func (this *HTTPServer) Handle(pattern string, handler http.Handler) {
	this.serveMux.Handle(pattern, handler)
}

func (this *HTTPServer) Start() (err error) {

	if this.listener, err = this.listen(this.Address); err != nil {
		return err
	}

	if this.listener != nil {
		if this.Static != "" && file.IsDir(this.Static) {
			var basepath = "/" + filepath.Base(this.Static) + "/"
			this.serveMux.Handle(basepath, http.StripPrefix(basepath, http.FileServer(http.Dir(this.Static))))
		}
		return this.server.Serve(this.listener)
	}

	return errors.New("listener empty")
}

func (this *HTTPServer) Stop() {

	//return this.server.Shutdown()
}

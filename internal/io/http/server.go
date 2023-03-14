package http

import (
	"fmt"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type Server struct {
	srv *fasthttp.Server
}

func New() *Server {
	r := router.New()
	registerRoutes(r)

	srv := &fasthttp.Server{
		Handler:                      r.Handler,
		ErrorHandler:                 nil,
		TCPKeepalive:                 true,
		DisablePreParseMultipartForm: true,
	}

	return &Server{
		srv: srv,
	}
}

func (s *Server) Run(port int) {
	go s.srv.ListenAndServe(fmt.Sprintf(":%d", port))
}

func (s *Server) Shutdown() error {
	return s.srv.Shutdown()
}

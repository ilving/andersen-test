package http

import (
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

func (s *Server) Run(host string) error {
	return s.srv.ListenAndServe(host)
}

func (s *Server) Shutdown() error {
	return s.srv.Shutdown()
}

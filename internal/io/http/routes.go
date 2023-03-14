package http

import (
	"awesomeProject/internal/io/http/handlers"
	"github.com/fasthttp/router"
)

func registerRoutes(r *router.Router) {
	r.Handle("POST", "/", handlers.FindPrimeNumbers)
}

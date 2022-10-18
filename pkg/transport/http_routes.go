package transport

import (
	"github.com/currency/pkg/currency/handler"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewHTTPRouter(cc handler.Handler) http.Handler {
	r := chi.NewMux()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/currency", cc.GetAllHandler)
	r.Get("/currency/{currency}", cc.GetByCurrencyHandler)

	return r
}

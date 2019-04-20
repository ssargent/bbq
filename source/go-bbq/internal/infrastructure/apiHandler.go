package infrastructure

import (
	"github.com/go-chi/chi"
)

type ApiHandler interface {
	Routes() *chi.Mux
}

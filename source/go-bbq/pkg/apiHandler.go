package pkg

import (
	"github.com/go-chi/chi"
)

type ApiHandler interface {
	Register() *chi.Mux
}

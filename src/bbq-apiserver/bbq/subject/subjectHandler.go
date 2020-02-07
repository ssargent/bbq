package subject

import (
	"net/http"

	//	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"github.com/google/uuid"

	//"github.com/google/uuid"

	"github.com/ssargent/bbq/bbq-apiserver/bbq"
	"github.com/ssargent/bbq/bbq-apiserver/config"
	"github.com/ssargent/bbq/bbq-apiserver/internal/infrastructure"
	"github.com/ssargent/bbq/bbq-apiserver/security"
)

type subjectHandler struct {
	service        bbq.SubjectService
	authentication security.AuthenticationService
	config         *config.Config
}

// NewSubjectHandler will create an api Handler for a Subjects.
func NewSubjectHandler(config *config.Config, authentication security.AuthenticationService, service bbq.SubjectService) infrastructure.ApiHandler {
	return &subjectHandler{service: service, authentication: authentication, config: config}
}

func (handler *subjectHandler) Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", handler.getSubjects)

	return router
}

func (handler *subjectHandler) getSubjects(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())

	tenantString := claims["tenant"].(string)
	tenant, err := uuid.Parse(tenantString)

	if err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}
	subjectData, err := handler.service.GetSubjects(tenant)

	if err != nil {
		render.Render(w, r, infrastructure.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, subjectData)
}

package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

type APIRoute http.Handler

type API struct {
	r *chi.Mux
}

type APIResponse struct {
	Status bool   `json:"status"`
	Error  string `json:"error,omitempty"`
}

func (api *API) AddRouter(route string, router APIRoute) {
	api.r.Mount(route, router)
}

func (api *API) Start(port string) error {
	log.Infof("[API] Listening for REST API Requests on port %v\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), api.r)
}

func (api *API) StartTLS(addr, certFile, keyFile string) error {
	return http.ListenAndServe(addr, certFile, keyFile, api.r)
}

func NewAPI() *API {
	r := chi.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrgins:   []string("*"),
		AllowOriginFunc: func(r *http.Request, orgin string) bool { return true },
	})

	r.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.AllowContentType("application/json"),
		middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: log.StandardLogger(), NoColor: true}),
		middleware.Compress(6, "gzip"),
		middleware.RedirectSlashes,
		middleware.Recoverer,
		middleware.NoCache,
		middleware.Timeout(60*time.Second),
		c.Handler,
	)

	return &API{r: r}
}

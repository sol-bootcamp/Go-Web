package server

import (
	"net/http"
	"web/internal/handler"
	mw "web/internal/middleware"
	"web/internal/repository"
	"web/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type ConfigServerChi struct {
	// ServerAddress is the address where the server will be listening
	ServerAddress string
	// LoaderFilePath is the path to the file that contains the vehicles
	LoaderFilePath string
}

// NewServerChi is a function that returns a new instance of ServerChi
func NewServerChi(cfg *ConfigServerChi) *ServerChi {
	// default values
	defaultConfig := &ConfigServerChi{
		ServerAddress: ":8080",
	}
	if cfg != nil {
		if cfg.ServerAddress != "" {
			defaultConfig.ServerAddress = cfg.ServerAddress
		}
		if cfg.LoaderFilePath != "" {
			defaultConfig.LoaderFilePath = cfg.LoaderFilePath
		}
	}

	return &ServerChi{
		serverAddress:  defaultConfig.ServerAddress,
		loaderFilePath: defaultConfig.LoaderFilePath,
	}
}

// ServerChi is a struct that implements the Application interface
type ServerChi struct {
	// serverAddress is the address where the server will be listening
	serverAddress string
	// loaderFilePath is the path to the file that contains the vehicles
	loaderFilePath string
}

// Run is a method that runs the application
func (a *ServerChi) Run(apiToken string) (err error) {
	// - repository
	pr, err := repository.NewProductRepository(a.loaderFilePath)
	if err != nil {
		return
	}
	// - service
	ps := service.NewProductService(pr)
	// - handler
	ph := handler.NewProductHandler(ps)

	// router
	r := chi.NewRouter()

	// - middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(mw.LoggingMiddleware)

	// - endpoints

	// Health check endpoint
	r.Group(func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong"))
		})
	})

	r.Route("/products", func(r chi.Router) {

		// Public endpoints
		r.Group(func(r chi.Router) {
			r.Get("/", ph.GetAllProducts())

			r.Get("/{id}", ph.GetProductByID)

			r.Get("/search", ph.SearchProduct)

		})
		//Private endpoints
		r.Group(func(r chi.Router) {
			// r.Use(Auth)
			r.Post("/", ph.CreateProduct)

			r.With(mw.Auth(apiToken)).Put("/{id}", ph.UpdateCreate)

			r.Patch("/{id}", ph.Patch)

			r.Delete("/", ph.DeleteProduct)
		})

	})

	// run server
	err = http.ListenAndServe(a.serverAddress, r)
	return
}

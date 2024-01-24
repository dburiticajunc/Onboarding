package application

import (
	"github.com/go-chi/chi/v5"
	"ms-product/internal/handler"
	"ms-product/internal/repository"
	"ms-product/internal/service"
	"ms-product/internal/storage"
	"net/http"
)

type ConfigAppDefault struct {
	ServerAddr string
	DbFile     string
}

type ApplicationDefault struct {
	rt         *chi.Mux
	serverAddr string
	DbFile     string
}

func NewApplicationDefault(cfg *ConfigAppDefault) *ApplicationDefault {
	defaultRouter := chi.NewRouter()
	defaultConfig := &ConfigAppDefault{
		ServerAddr: ":8080",
		DbFile:     "",
	}
	if cfg != nil {
		if cfg.ServerAddr != "" {
			defaultConfig.ServerAddr = cfg.ServerAddr
		}
		if cfg.DbFile != "" {
			defaultConfig.DbFile = cfg.DbFile
		}
	}
	return &ApplicationDefault{
		rt:         defaultRouter,
		serverAddr: defaultConfig.ServerAddr,
		DbFile:     defaultConfig.DbFile,
	}
}

func (a *ApplicationDefault) SetUp() (err error) {
	//dependencies
	db, err := storage.NewStorageProductJSON(a.DbFile).ReadAll()
	if err != nil {
		return
	}

	// initializer dependencies
	rp := repository.NewRepositoryProductMap(len(db), db)
	sv := service.NewServiceProductDefault(rp)
	hd := handler.NewHandlerProductDefault(sv)

	// initializer router
	(*a).rt.Route("/product", func(r chi.Router) {
		r.Get("/health", func(wr http.ResponseWriter, re *http.Request) {
			wr.WriteHeader(http.StatusOK)
			wr.Header().Set("Content-Type", "text/plain")
			wr.Write([]byte("UP"))
		})
		r.Get("/getAll", hd.GetAll())
	})
	return
}

func (a *ApplicationDefault) Run() (err error) {
	err = http.ListenAndServe(a.serverAddr, a.rt)
	return
}

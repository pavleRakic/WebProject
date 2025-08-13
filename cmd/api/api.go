package api

import (
	"database/sql"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	permission "github.com/pavleRakic/testGoApi/service/permission"
	product "github.com/pavleRakic/testGoApi/service/quiz"
	resource "github.com/pavleRakic/testGoApi/service/resource"
	"github.com/pavleRakic/testGoApi/service/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()

	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(subrouter)

	resourceStore := resource.NewStore(s.db)
	resourceHandler := resource.NewHandler(resourceStore)
	resourceHandler.RegisterRoutes(subrouter)

	permissionStore := permission.NewStore(s.db)
	permissionHandler := permission.NewHandler(permissionStore)
	permissionHandler.RegisterRoutes(subrouter)

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./cmd/static/"))))

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		absPath, err := filepath.Abs("./static/main.html")
		if err != nil {
			log.Println("Error getting absolute path:", err)
		} else {
			log.Println("Serving file from:", absPath)
		}
		http.ServeFile(w, r, "./cmd/static/main.html")

	})

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("NOT FOUND:", r.URL.Path)
		http.NotFound(w, r)
	})

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}

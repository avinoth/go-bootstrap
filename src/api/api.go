package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

var routerInstance *negroni.Negroni

func initializeRoutes(router *mux.Router) {
	router.HandleFunc("/ping", Ping).Methods("GET")
}

func setContentType(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Header().Set("Content-Type", "application/json")
	next(w, r)
}

// InitializeRouter initializes the routes, sets middleware and returns a negroni instance.
func InitializeRouter() *negroni.Negroni {
	router := mux.NewRouter()
	initializeRoutes(router)

	n := negroni.New()

	n.Use(negroni.HandlerFunc(setContentType))
	n.UseHandler(router)
	routerInstance = n
	return n
}

// RunServer accepts port and starts the web server
func RunServer(port string) {
	InitializeRouter()
	routerInstance.Run(":" + port)
}

// Instance returns the currently initialized negroni instance
func Instance() *negroni.Negroni {
	return routerInstance
}

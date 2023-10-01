package app

import (
	"github.com/gorilla/mux"
	"sewa-in/app/controllers"
)

func (server *Server) initializeRoutes() {
	server.Router = mux.NewRouter()

	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}

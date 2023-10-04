package app

import (
	"github.com/gorilla/mux"
	"sewa-in/app/controllers"
)

func (server *Server) initializeRoutes() {
	server.Router = mux.NewRouter()

	//sewa.in - internal
	//GPS_tracer
	server.Router.HandleFunc("/tracer/get", controllers.Home).Methods("POST")
	server.Router.HandleFunc("/tracer/set/interval", controllers.Home).Methods("POST")

	//sewa.in - external
	//user
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
	server.Router.HandleFunc("/home", controllers.Home).Methods("GET")
	server.Router.HandleFunc("/register", controllers.Home).Methods("POST")
	server.Router.HandleFunc("/login", controllers.Home).Methods("GET")

	//stockb
	//server.Router.HandleFunc("/stock", controllers.Home).Methods("GET")
	server.Router.HandleFunc("/stock/create", controllers.Home).Methods("POST")
	server.Router.HandleFunc("/stock/detail", controllers.Home).Methods("GET")
	server.Router.HandleFunc("/stock/list", controllers.Home).Methods("GET")
	server.Router.HandleFunc("/stock/update", controllers.Home).Methods("PUT")

	//transaction
	server.Router.HandleFunc("/transaction", controllers.Home).Methods("POST")
	server.Router.HandleFunc("/transaction/list", controllers.Home).Methods("GET")
	server.Router.HandleFunc("/transaction/detail", controllers.Home).Methods("GET")
	server.Router.HandleFunc("/transaction/active", controllers.Home).Methods("GET")
	server.Router.HandleFunc("/transaction/finised", controllers.Home).Methods("GET")
	server.Router.HandleFunc("/transaction/review", controllers.Home).Methods("POST")

}

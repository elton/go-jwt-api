package controllers

import "github.com/elton/go-jwt-api/api/middlewares"

func (server *Server) initializeRoutes() {
	// Home Route
	server.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(server.Home)).Methods("GET")

	// Login Route
	server.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(server.Login)).Methods("POST")

	//Users routes
	server.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.CreateUser)).Methods("POST")
	server.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.GetUsers)).Methods("GET")
	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(server.GetUser)).Methods("GET")
	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareJSON(server.UpdateUser))).Methods("PUT")
	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(server.DeleteUser)).Methods("DELETE")
}

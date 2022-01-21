package controllers

import "github.com/igorariza/Dockerized-Golang_API-MySql-React.js/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	//Users routes

	//s.Router.HandleFunc("/api/v1/team", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/api/v1/players", middlewares.SetMiddlewareJSON(s.GetPlayers)).Methods("GET")
	//s.Router.HandleFunc("/api/v1/player", middlewares.SetMiddlewareJSON(s.GetAllUsers)).Methods("GET")

}

package controllers

import (
	"net/http"

	"github.com/igorariza/Dockerized-Golang_API-MySql-React.js/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}

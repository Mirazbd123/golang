package router

import (
	"github.com/gorilla/mux"
	"github.com/mirazbd123/mongoapi/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movies", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movies/{id}", controller.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("api/deleteallmovies", controller.DeleteAllMovie).Methods("DELETE")

	return router
}

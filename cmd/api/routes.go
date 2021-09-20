package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	router.HandlerFunc(http.MethodGet, "/api/v1/movies/:id", app.getOneMovie)
	router.HandlerFunc(http.MethodGet, "/api/v1/movies", app.getAllMovie)
	router.HandlerFunc(http.MethodGet, "/api/v1/genres/:genre_id", app.getAllMoviesByGenre)

	router.HandlerFunc(http.MethodGet, "/api/v1/genres", app.getAllGenres)

	return app.enableCORS(router)
}

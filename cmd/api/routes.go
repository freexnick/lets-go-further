package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)
	mux.HandleFunc("/v1/movies", app.listMoviesHandler)
	mux.HandleFunc("/v1/movies/{id}", app.showMovieHandler)
	mux.HandleFunc("POST /v1/movies", app.createMovieHandler)
	mux.HandleFunc("PATCH /v1/movies/{id}", app.updateMovieHandler)
	mux.HandleFunc("DELETE /v1/movies/{id}", app.deleteMovieHandler)

	mux.HandleFunc("POST /v1/users", app.registerUserHandler)

	return app.recoverPanic(app.rateLimit(mux))
}

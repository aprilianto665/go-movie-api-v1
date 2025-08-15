package routes

import (
	"movie-api-v1/controllers"
	"net/http"
)

func MovieRouter() *http.ServeMux {
	mux := http.NewServeMux()
	
	mux.HandleFunc("GET /movies", controllers.GetAllMovies)
	mux.HandleFunc("GET /movie/{id}", controllers.GetMovieById)
	mux.HandleFunc("PUT /movie/{id}", controllers.UpdateMovieById)
	mux.HandleFunc("DELETE /movie/{id}",controllers.DeleteMovieById)
	mux.HandleFunc("POST /movie", controllers.CreateMovie)

	return mux
}
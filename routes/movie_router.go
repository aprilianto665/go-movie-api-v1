package routes

import (
	"movie-api-v1/controllers"
	"net/http"
)

func MovieRouter(controller *controllers.MovieController) *http.ServeMux {
	mux := http.NewServeMux()
	
	mux.HandleFunc("GET /movies", controller.GetAllMovies)
	mux.HandleFunc("GET /movie/{id}", controller.GetMovieById)
	mux.HandleFunc("PUT /movie/{id}", controller.UpdateMovieById)
	mux.HandleFunc("DELETE /movie/{id}", controller.DeleteMovieById)
	mux.HandleFunc("POST /movie", controller.CreateMovie)

	return mux
}
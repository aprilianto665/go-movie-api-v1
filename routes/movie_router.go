package routes

import (
	"database/sql"
	"movie-api-v1/controllers"
	"net/http"
)

func MovieRouter(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	h := controllers.NewMovieController(db)
	
	mux.HandleFunc("GET /movies", h.GetAllMovies)
	mux.HandleFunc("GET /movie/{id}", h.GetMovieById)
	mux.HandleFunc("PUT /movie/{id}", h.UpdateMovieById)
	mux.HandleFunc("DELETE /movie/{id}",h.DeleteMovieById)
	mux.HandleFunc("POST /movie", h.CreateMovie)

	return mux
}
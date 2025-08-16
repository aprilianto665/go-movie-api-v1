package controllers

import (
	"encoding/json"
	"movie-api-v1/models"
	"movie-api-v1/repositories"
	"net/http"
	"strconv"
)

type MovieController struct {
    repo repositories.MovieRepository
}

func NewMovieController(repo repositories.MovieRepository) *MovieController {
    return &MovieController{repo: repo}
}

func (c *MovieController) GetAllMovies(w http.ResponseWriter, r *http.Request) {
    movies, err := c.repo.GetAll()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(movies)
}

func (c *MovieController) GetMovieById(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }
    
    movie, err := c.repo.GetByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(movie)
}

func (c *MovieController) CreateMovie(w http.ResponseWriter, r *http.Request) {
    var movie models.Movie
    if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    if err := c.repo.Create(&movie); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(movie)
}

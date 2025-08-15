package controllers

import (
	"database/sql"
	"fmt"
	"movie-api-v1/models"
	"net/http"
)



func NewMovieController(db *sql.DB) *models.MovieController { return &models.MovieController{db: db} }

func (c *models.MovieController) GetAllMovies (writer http.ResponseWriter, request *http.Request){
	rows, err := c.db.Query("SELECT * FROM movies")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
} 

func (c *models.MovieController) GetMovieById (writer http.ResponseWriter, request *http.Request){
	id := request.PathValue("id")
	fmt.Fprintf(writer, "Get movie by ID %s", id)
} 

func UpdateMovieById (writer http.ResponseWriter, request *http.Request){
	fmt.Fprint(writer, "Update movie by ID")
} 

func DeleteMovieById (writer http.ResponseWriter, request *http.Request){
	fmt.Fprint(writer, "Delete movie by ID")
} 

func CreateMovie (writer http.ResponseWriter, request *http.Request){
	fmt.Fprint(writer, "Create movie")
} 
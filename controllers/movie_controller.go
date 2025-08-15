package controllers

import (
	"fmt"
	"net/http"
)

func GetAllMovies (writer http.ResponseWriter, request *http.Request){
	fmt.Fprint(writer, "Get all movies")
} 

func GetMovieById (writer http.ResponseWriter, request *http.Request){
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
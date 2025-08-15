package main

import (
	"log"
	"net/http"

	"movie-api-v1/db"
	"movie-api-v1/routes"
)

func main(){
	dsn := "postgres://ian:postgres@localhost:5432/movie_api?sslmode=disable"
	db, err := db.Connect(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mux := routes.MovieRouter(db)

	server := http.Server{
		Addr: ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
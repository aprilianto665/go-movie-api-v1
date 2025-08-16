package main

import (
	"log"
	"movie-api-v1/controllers"
	"movie-api-v1/db"
	"movie-api-v1/repositories"
	"movie-api-v1/routes"
	"net/http"
)

func main() {
    dsn := "postgres://user:password@localhost:5432/go_movie_v1?sslmode=disable"
    database, err := db.Connect(dsn)
    if err != nil {
        log.Fatal(err)
    }
    defer database.Close()

    movieRepo := repositories.NewMovieRepository(database)
    movieController := controllers.NewMovieController(movieRepo)
    
    mux := routes.MovieRouter(movieController)

    server := http.Server{
        Addr:    ":8080",
        Handler: mux,
    }

    log.Println("Server running on :8080")
    if err := server.ListenAndServe(); err != nil {
        panic(err)
    }
}

package main

import (
	"movie-api-v1/routes"
	"net/http"
)

func main(){
	mux := routes.MovieRouter()

	server := http.Server{
		Addr: ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
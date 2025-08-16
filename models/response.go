package models

type Response struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}

type MovieResponse struct {
    Status  string `json:"status"`
    Message string `json:"message"`
    Data    Movie  `json:"data"`
}

type MoviesResponse struct {
    Status  string  `json:"status"`
    Message string  `json:"message"`
    Data    []Movie `json:"data"`
}
package models

import "database/sql"

type MovieController struct{ db *sql.DB }

type Movie struct {
    ID    int64  `json:"id"`
    Title string `json:"title"`
    Year  int    `json:"year"`
}
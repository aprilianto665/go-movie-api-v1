package models

type Movie struct {
    ID          int      `json:"id"`
    Title       string   `json:"title"`
    Image       string   `json:"image"`
    Category    string   `json:"category"`
    Genre       []string `json:"genre"`
    Duration    int      `json:"duration"`
    ReleaseYear int      `json:"releaseYear"`
    AgeRating   int      `json:"ageRating"`
    Synopsis    string   `json:"synopsis"`
    Cast        []string `json:"cast"`
    Creators    []string `json:"creators"`
    Rating      float64  `json:"rating"`
}
// repositories/movie_repository_impl.go
package repositories

import (
	"database/sql"
	"movie-api-v1/models"

	"github.com/lib/pq"
)

type movieRepository struct {
    db *sql.DB
}

func NewMovieRepository(db *sql.DB) MovieRepository {
    return &movieRepository{db: db}
}

func (r *movieRepository) GetAll() ([]models.Movie, error) {
    query := `SELECT id, title, image, category, genre, duration, release_year, age_rating, synopsis, cast, creators, rating FROM movies`
    
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var movies []models.Movie
    for rows.Next() {
        var movie models.Movie
        err := rows.Scan(
            &movie.ID, &movie.Title, &movie.Image, &movie.Category,
            pq.Array(&movie.Genre), &movie.Duration, &movie.ReleaseYear,
            &movie.AgeRating, &movie.Synopsis, pq.Array(&movie.Cast),
            pq.Array(&movie.Creators), &movie.Rating,
        )
        if err != nil {
            return nil, err
        }
        movies = append(movies, movie)
    }
    return movies, nil
}

func (r *movieRepository) GetByID(id int) (*models.Movie, error) {
    query := `SELECT id, title, image, category, genre, duration, release_year, age_rating, synopsis, cast, creators, rating FROM movies WHERE id = $1`
    
    var movie models.Movie
    err := r.db.QueryRow(query, id).Scan(
        &movie.ID, &movie.Title, &movie.Image, &movie.Category,
        pq.Array(&movie.Genre), &movie.Duration, &movie.ReleaseYear,
        &movie.AgeRating, &movie.Synopsis, pq.Array(&movie.Cast),
        pq.Array(&movie.Creators), &movie.Rating,
    )
    if err != nil {
        return nil, err
    }
    return &movie, nil
}

func (r *movieRepository) Create(movie *models.Movie) error {
    query := `INSERT INTO movies (title, image, category, genre, duration, release_year, age_rating, synopsis, cast, creators, rating) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id`
    
    err := r.db.QueryRow(query, movie.Title, movie.Image, movie.Category,
        pq.Array(movie.Genre), movie.Duration, movie.ReleaseYear,
        movie.AgeRating, movie.Synopsis, pq.Array(movie.Cast),
        pq.Array(movie.Creators), movie.Rating).Scan(&movie.ID)
    
    return err
}

func (r *movieRepository) Update(id int, movie *models.Movie) error {
    query := `UPDATE movies SET title=$1, image=$2, category=$3, genre=$4, duration=$5, 
              release_year=$6, age_rating=$7, synopsis=$8, cast=$9, creators=$10, rating=$11 
              WHERE id=$12`
    
    _, err := r.db.Exec(query, movie.Title, movie.Image, movie.Category,
        pq.Array(movie.Genre), movie.Duration, movie.ReleaseYear,
        movie.AgeRating, movie.Synopsis, pq.Array(movie.Cast),
        pq.Array(movie.Creators), movie.Rating, id)
    
    return err
}

func (r *movieRepository) Delete(id int) error {
    query := `DELETE FROM movies WHERE id = $1`
    _, err := r.db.Exec(query, id)
    return err
}

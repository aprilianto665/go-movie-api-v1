package repositories

import (
	"movie-api-v1/models"
)

type MovieRepository interface {
    GetAll() ([]models.Movie, error)
    GetByID(id int) (*models.Movie, error)
    Create(movie *models.Movie) error
    Update(id int, movie *models.Movie) error
    Delete(id int) error
}
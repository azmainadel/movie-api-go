package repository

import (
	"github.com/azmainadel/movie-api-go/model"
)

type IMovieRepository interface {
	GetAllMovies() ([]model.Movie, error)
	GetMovie(id string) (model.Movie, error)
	CreateMovie(movie model.Movie) error
	UpdateMovie(id string, movie model.Movie) error
	DeleteMovie(id string) error
	DeleteAllMovies() error
}

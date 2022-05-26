package service

import "github.com/azmainadel/movie-api-go/model"

type IMovieService interface {
	GetAllMovies() ([]model.Movie, error)
	GetMovie(id int) (model.Movie, error)
	CreateMovie(movie model.Movie) error
	UpdateMovie(id int, movie model.Movie) error
	DeleteMovie(id int) error
	DeleteAllMovies() error
}

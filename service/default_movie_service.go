package service

import (
	"errors"
	"github.com/azmainadel/movie-api-go/model"
	"github.com/azmainadel/movie-api-go/repository"
	"github.com/azmainadel/movie-api-go/utility"
)

type DefaultMovieService struct {
	movieRepository repository.IMovieRepository
}

func NewDefaultMovieService(mRepo repository.IMovieRepository) *DefaultMovieService {
	return &DefaultMovieService{
		movieRepository: mRepo,
	}
}

func (d *DefaultMovieService) GetAllMovies() ([]model.Movie, error) {
	return d.movieRepository.GetAllMovies()
}

func (d *DefaultMovieService) GetMovie(id string) (model.Movie, error) {
	movie, err := d.movieRepository.GetMovie(id)

	if err != nil {
		if errors.Is(err, utility.ErrMovieNotFound) {
			return model.Movie{}, utility.ErrMovieNotFound
		}
	}
	return movie, nil
}

func (d *DefaultMovieService) CreateMovie(movie model.Movie) error {
	if movie.Title == "" {
		return utility.ErrInvalidTitleField
	}
	return d.movieRepository.CreateMovie(movie)
}

func (d *DefaultMovieService) UpdateMovie(id string, movie model.Movie) error {
	if movie.Title == "" {
		return utility.ErrInvalidTitleField
	}

	err := d.movieRepository.UpdateMovie(id, movie)

	if errors.Is(err, utility.ErrMovieNotFound) {
		return utility.ErrMovieNotFound
	}

	return nil
}

func (d *DefaultMovieService) DeleteMovie(id string) error {
	err := d.movieRepository.DeleteMovie(id)

	if err != nil {
		if errors.Is(err, utility.ErrMovieNotFound) {
			return utility.ErrMovieNotFound
		}
		return err
	}

	return nil
}

func (d *DefaultMovieService) DeleteAllMovies() error {
	return d.movieRepository.DeleteAllMovies()
}

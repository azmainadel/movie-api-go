package repository

import (
	"github.com/azmainadel/movie-api-go/model"
	"github.com/azmainadel/movie-api-go/utility"
	"github.com/google/uuid"
)

type localMovieRepository struct {
	Movies []model.Movie
}

func NewLocalMovieRepository() *localMovieRepository {
	var movies = []model.Movie{
		{Id: uuid.NewString(), Title: "The Shawshank Redemption", ReleaseYear: 1994, Director: "Frank Darabont", Score: 9.3},
		{Id: uuid.NewString(), Title: "The Godfather", ReleaseYear: 1972, Director: "Francis Ford Coppola", Score: 9.2},
		{Id: uuid.NewString(), Title: "The Dark Knight", ReleaseYear: 2008, Director: "Christopher Nolan", Score: 9.0},
	}

	return &localMovieRepository{
		Movies: movies,
	}
}

func (i *localMovieRepository) GetAllMovies() ([]model.Movie, error) {
	return i.Movies, nil
}

func (i *localMovieRepository) GetMovie(id string) (model.Movie, error) {
	for _, movie := range i.Movies {
		if movie.Id == id {
			return movie, nil
		}
	}

	return model.Movie{}, utility.ErrMovieNotFound
}

func (i *localMovieRepository) CreateMovie(movie model.Movie) error {
	movie.Id = uuid.NewString()
	i.Movies = append(i.Movies, movie)

	return nil
}

func (i *localMovieRepository) UpdateMovie(id string, movie model.Movie) error {
	for _, m := range i.Movies {
		if m.Id == id {
			m.Title = movie.Title
			m.ReleaseYear = movie.ReleaseYear
			m.Director = movie.Director
			m.Score = movie.Score

			return nil
		}
	}

	return utility.ErrMovieNotFound
}

func (i *localMovieRepository) DeleteMovie(id string) error {
	movieFound := false

	var updatedMovieList []model.Movie

	for _, movie := range i.Movies {
		if movie.Id == id {
			movieFound = true
		} else {
			updatedMovieList = append(updatedMovieList, movie)
		}
	}

	if !movieFound {
		return utility.ErrMovieNotFound
	}

	i.Movies = updatedMovieList

	return nil
}

func (i *localMovieRepository) DeleteAllMovies() error {
	i.Movies = nil

	return nil
}

package handler

import (
	"encoding/json"
	"errors"
	"github.com/azmainadel/movie-api-go/model"
	"github.com/azmainadel/movie-api-go/service"
	"github.com/azmainadel/movie-api-go/utility"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type movieHandler struct {
	service service.IMovieService
}

func NewMovieHandler(mServ service.IMovieService) *movieHandler {
	return &movieHandler{service: mServ}
}

func (mHandler *movieHandler) GetAllMovies(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	movies, err := mHandler.service.GetAllMovies()

	if err != nil {
		http.Error(w, "Movie fetching error", http.StatusInternalServerError)
		return
	}

	movieJson, err := json.Marshal(movies)

	if err != nil {
		http.Error(w, "Movie data parsing error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(movieJson)
}

func (mHandler *movieHandler) GetMovie(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	movie, err := mHandler.service.GetMovie(ps.ByName("id"))

	if err != nil {
		if errors.Is(err, utility.ErrInvalidID) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else if errors.Is(err, utility.ErrMovieNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	movieJson, err := json.Marshal(movie)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(movieJson)
}

func (mHandler *movieHandler) CreateMovie(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var movie model.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)

	if err != nil {
		http.Error(w, "Movie data JSON decoding error", http.StatusInternalServerError)
		return
	}

	err = mHandler.service.CreateMovie(movie)

	if err != nil {
		if errors.Is(err, utility.ErrInvalidTitleField) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("Movie stored in database successfully"))
}

func (mHandler *movieHandler) UpdateMovie(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var movie model.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)

	if err != nil {
		http.Error(w, "Movie data JSON decoding error", http.StatusInternalServerError)
		return
	}

	err = mHandler.service.UpdateMovie(ps.ByName("id"), movie)

	if err != nil {
		if errors.Is(err, utility.ErrInvalidID) || errors.Is(err, utility.ErrInvalidTitleField) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else if errors.Is(err, utility.ErrMovieNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (mHandler *movieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := mHandler.service.DeleteMovie(ps.ByName("id"))

	if err != nil {
		if errors.Is(err, utility.ErrInvalidID) || errors.Is(err, utility.ErrInvalidTitleField) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
	_, _ = w.Write([]byte("Movie data deleted from database successfully"))
}

func (mHandler *movieHandler) DeleteAllMovies(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := mHandler.service.DeleteAllMovies()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	_, _ = w.Write([]byte("All movie data deleted from database successfully"))
}

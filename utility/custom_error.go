package utility

import "errors"

var (
	ErrInvalidID         = errors.New("invalid ID provided")
	ErrMovieNotFound     = errors.New("movie not found")
	ErrInvalidTitleField = errors.New("title field cannot be empty")
)

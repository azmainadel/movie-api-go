package main

import (
	"github.com/azmainadel/movie-api-go/handler"
	"github.com/azmainadel/movie-api-go/repository"
	"github.com/azmainadel/movie-api-go/service"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	movieLocalRepository := repository.NewLocalMovieRepository()
	movieService := service.NewDefaultMovieService(movieLocalRepository)
	movieHandler := handler.NewMovieHandler(movieService)

	router := httprouter.New()

	router.GET("/movies", movieHandler.GetAllMovies)
	router.GET("/movies/:id", movieHandler.GetMovie)

	router.POST("/create-movie", movieHandler.CreateMovie)

	router.PATCH("/update-movie", movieHandler.UpdateMovie)

	router.DELETE("/delete-movies", movieHandler.DeleteAllMovies)
	router.PATCH("/delete-movies/:id", movieHandler.DeleteMovie)

	log.Println("HTTP server runs on :8080")
	err := http.ListenAndServe(":8080", router)

	log.Fatal(err)
}

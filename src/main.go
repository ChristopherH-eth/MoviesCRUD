package main

/**
 * @file main.go
 * @author Original author: Free Code Camp
 *		   Changes made by 0xChristopher for learning purposes
 *
 * A simple CRUD implementation which starts up a webserver and allows user to view all of the stored movies,
 * view movies by id, update movie entries, as well as delete movies.
 */

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

// The GetMovies() functions returns the entire 'movies' slice.
func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(movies)
}

// The DeleteMovie() function deletes a movie from the 'movies' slice by id.
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	params := mux.Vars(r)

	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)
}

// The GetMovie() function returns a movie from the 'movies' slice by id.
func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	params := mux.Vars(r)

	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode((item))
			return
		}
	}
}

// The CreateMovie() function creates a new movie and appends it to the 'movies' slice.
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Aplication/json")

	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

// The UpdateMovie() function updates a movie entry by removing the 'movie' at the id and adding a new
// 'movie' at that same id.
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	params := mux.Vars(r)

	for index, item := range movies {
		if item.Id == params["id"] {
			// Remove movie
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)

			// Add movie
			movie.Id = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	// Create initial movies
	movies = append(movies, Movie{Id: "1", Isbn: "438227", Title: "Movie 1", Director: &Director{
		Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{Id: "2", Isbn: "45455", Title: "Movie 2", Director: &Director{
		Firstname: "Steve", Lastname: "Smith"}})

	// Define routes
	r.HandleFunc("/movies", GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", GetMovie).Methods("GET")
	r.HandleFunc("/movies", CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", DeleteMovie).Methods("DELETE")

	// Server startup
	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}

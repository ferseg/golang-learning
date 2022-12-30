package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request)  {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(movies)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    newMovie := Movie{}
    json.NewDecoder(r.Body).Decode(&newMovie)
    newMovie.Id = strconv.Itoa(rand.Int())
    movies = append(movies, newMovie)
    json.NewEncoder(w).Encode(newMovie)
}

func deleteMovie(w http.ResponseWriter, r * http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id := params["id"]
    for index, item := range movies {
        if item.Id == id {
            movies = append(movies[:index], movies[index + 1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r * http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id := params["id"]
    for _, item := range movies {
        if item.Id == id {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    // fmt.Fprintf("Movie not found %v", http.StatusNotFound)
}

func updateMovie(w http.ResponseWriter, r * http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id := params["id"]
    updatedMovie := Movie{}
    json.NewDecoder(r.Body).Decode(&updatedMovie)
    for index, item := range movies {
        if item.Id == id {
            movies = append(movies[:index], movies[index + 1:]...)
            updatedMovie.Id = id
            movies = append(movies, updatedMovie)
            json.NewEncoder(w).Encode(updatedMovie)
            break
        }
    }
}

func main() {
    r := mux.NewRouter()

    starWars := Movie{
        Id: "100",
        Isbn: "33254",
        Title: "Start Wars, Clone Wars",
        Director: &Director{
            FirstName: "George",
            LastName: "Lucas",
        },
    }

    movies = append(movies, starWars)

    // /movies
    r.HandleFunc("/movies", getMovies).Methods("GET")
    r.HandleFunc("/movies", createMovie).Methods("POST")

    // /movies/{id}
    r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
    r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
    r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

    fmt.Println("Starting server on port 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}

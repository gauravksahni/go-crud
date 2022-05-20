package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

// 
// 
type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"titile"`
	Director *Director `json:"director"`

}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"Lastname"`
}

var movies []Movie



func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// func deleteMovie(w http.ResponseWriter, r *http.Request){
// 	w.Header().Set("Content-Type", "application/json")
// }

func main(){
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isbn: "4321", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "5321", Title: "Movie Two", Director: &Director{Firstname: "Deen", Lastname: "Doe"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	// r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	// r.HandleFunc("/movies", createMovie).Methods("POST")
	// r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	// r.HandleFunc("movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting the server at aport 8000 \n")
	log.Fatal(http.ListenAndServe(":8000",r))
		
}
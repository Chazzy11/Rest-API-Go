// creating a basic server - main is the entry point of the program

package main

import (
	"fmt" // package for formatting and printing
	"net/http" // package for http based web programs
	"encoding/json" // package for encoding and decoding json
	"log" // package for logging errors
	"github.com/gorilla/mux" // package for router from gorilla framework
)

type song struct {
	ID string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Artist string `json:"artist,omitempty"`
} // struct for song object

var songs []song // slice for songs

func getSongs(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(songs)
} // function to get all songs

func getSong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // get params from request body and store in params variable 
	for_, item := range songs { // loop through songs
		if item.ID == params["id"] { // if song id matches the id in params
			json.NewEncoder(w).Encode(item) // encode the song and return
			return
		}
	}
	json.NewEncoder(w).Encode(&song{}) // if no song found return empty song
} // function to get a song

func createSong(w http.ResponseWriter, r *http.Request) {
	var song song // create a song object
	_ = json.NewDecoder(r.Body).Decode(&song) // decode the request body and store in song object
	songs = append(songs, song) // append the song to the songs slice
	json.NewEncoder(w).Encode(song) // encode the song and return



func main() {
	router:= mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World") // handler function for the / route
	}	)

	http.ListenAndServe(":8080", router)
}


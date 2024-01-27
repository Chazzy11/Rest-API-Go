// creating a basic server - main is the entry point of the program

package main

import (
	// "fmt" // package for formatting and printing
	"net/http" // package for http based web programs
	"encoding/json" // package for encoding and decoding json
	"log" // package for logging errors
	"github.com/gorilla/mux" // package for router from gorilla framework
	"github.com/getsentry/sentry-go" // package for error tracking
	"time" // package for time based functions
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
	for _, item := range songs { // loop through songs, underscore is used to ignore the index
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
} // function to create a song

func updateSong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // get params from request body and store in params variable
	for index, item := range songs { // loop through songs
		if item.ID == params["id"] { // if song id matches the id in params
			songs = append(songs[:index], songs[index+1:]...) // delete the song
			var song song // create a song object
			_ = json.NewDecoder(r.Body).Decode(&song) // decode the request body and store in song object
			song.ID = params["id"] // assign the song id to the id in params
			songs = append(songs, song) // append the song to the songs slice
			json.NewEncoder(w).Encode(song) // encode the song and return
			return
		}
	}
	json.NewEncoder(w).Encode(songs) // if no song found return all songs
} // function to update a song

func deleteSong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // get params from request body and store in params variable
	for index, item := range songs { // loop through songs
		if item.ID == params["id"] { // if song id matches the id in params
			songs = append(songs[:index], songs[index+1:]...) // delete the song
			break
		}
	}
	json.NewEncoder(w).Encode(songs) // return all songs
} // function to delete a song


func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://4a9acdd1534746f0eb5ac6132fd8cadd@o4506638842986496.ingest.sentry.io/4506638866120707",
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	  })
	  if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	  }
	  // Flush buffered events before the program terminates.
	  defer sentry.Flush(2 * time.Second)
	
	  sentry.CaptureMessage("It works!")
	
	router:= mux.NewRouter()

	songs = append(songs, song{ID: "1", Title: "Lose Control", Artist: "Evanescence"})
	songs = append(songs, song{ID: "2", Title: "Amarylis", Artist: "Shinedown"})

	router.HandleFunc("/songs", getSongs).Methods("GET")
	router.HandleFunc("/songs/{id}", getSong).Methods("GET")
	router.HandleFunc("/songs", createSong).Methods("POST")
	router.HandleFunc("/songs/{id}", updateSong).Methods("PUT")
	router.HandleFunc("/songs/{id}", deleteSong).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}


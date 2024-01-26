// creating a basic server - main is the entry point of the program

package main

import (
	"fmt" // package for formatting and printing
	"net/http" // package for http based web programs

	"github.com/gorilla/mux" // package for router from gorilla framework
)

func main() {
	router:= mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World") // handler function for the / route
	}	)

	http.ListenAndServe(":8080", router)
}
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// r.HandleFunc("/", serveHome)
	r.HandleFunc("/", serveHome).Methods("GET") // tells that we are only accepting GET on this

	// http.ListenAndServe(":4000", r) // to start up a server in go. pass the port and the router
	// we can either handle above using comma ok syntax but another way is
	log.Fatal(http.ListenAndServe(":4000", r))
	// Fatal is equivalent to Print followed by a call to os.Exit(1).

	// Now run main.go and on browser on localhost:4000 you will be able to see the output
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	// params etc are in the request
	// what we have to send are in the response.

	w.Write([]byte("<h1>Golang YT</h1>"))
}

package main

import (
	handler "handler/Handler"
	"log"
	"net/http"
)

// Handler is a interface
// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }

// HanlderFunc allows normal functions to work for http handlers

func main() {
	mux := defaultMux()

	mp := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := handler.MapHandler(mp, mux)

	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`

	yampHandler, err := handler.YAMLHandler([]byte(yaml), mapHandler)
	CheckErrNil(err)

	log.Fatal(http.ListenAndServe(":8080", yampHandler))
}

func defaultMux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ye link daala hai map me pehle ?"))
}

func CheckErrNil(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	// "html/template"
	"net/http"

	"log"
	"os"
)

func main() {
	story := Story{}
	data, err := os.ReadFile("gopher.json")
	CheckErrNil(err)
	json.Unmarshal(data, &story)

	port := 8080
	// tpl := template.Must(template.New("gg").Parse("heloo"))
	// handler := NewHandler(story, WithTemplates(tpl))
	handler := NewHandler(story)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}

func CheckErrNil(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// mux by http.NewServeMux basically itself returns 404 if not mapped to handler instead of we handling

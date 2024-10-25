package handler

import (
	// "fmt"

	"net/http"

	"gopkg.in/yaml.v2"
	// "github.com/gorilla/mux"
	// "gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if l, found := pathsToUrls[r.URL.Path]; found {
			w.Header().Add("Content-Type", "application/json")
			// the code in below is necessary and not 3xx but proper http. ...
			http.Redirect(w, r, l, http.StatusMovedPermanently)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

type pathUrl struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var arr []pathUrl
	err := yaml.Unmarshal(yml, &arr)
	if err != nil {
		return nil, err
	}
	mp := make(map[string]string)
	for _, pU := range arr {
		mp[pU.Path] = pU.Url
	}
	return MapHandler(mp, fallback), nil
}

package main

import (
	"html/template"
	"net/http"
	"strings"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

var defaultHandlerTmpl = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <section class="page">
      <h1>{{.Title}}</h1>
      {{range .Paragraphs}}
        <p>{{.}}</p>
      {{end}}
      {{if .Options}}
        <ul>
        {{range .Options}}
          <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
        {{end}}
        </ul>
      {{else}}
        <h3>The End</h3>
      {{end}}
    </section>
    <style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FFFCF6;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #777;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: none;
        color: #6295b5;
      }
      a:active,
      a:hover {
        color: #7792a2;
      }
      p {
        text-indent: 1em;
      }
    </style>
  </body>
</html>`

// So init is a function which runs when your service is being up. See
// TODO  -   https://tutorialedge.net/golang/the-go-init-function/

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTmpl))
	// must as if template doesn't gets compiled correctly then our code is really wrong and doesn't make sense to error handle
}

// https://pkg.go.dev/text/template@go1.23.0 - nice self explanatory
var tpl *template.Template

// A good practice that if you are not exporting a type (means private) then better to return its interface rather than the specific type. This way end user will be able to see what methods that type has else won't be able to see.

func NewHandler(s Story, options ...HandlerOption) http.Handler {
	h := handler{s, tpl}
	for _, opt := range options {
		opt(&h)
	}

	return h
}

type handler struct {
	s Story
	t *template.Template
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]
	if value, found := h.s[path]; found {
		err := h.t.Execute(w, value)
		if err != nil {
			http.Error(w, "Sorry Something Went Wrong", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Not found Story", http.StatusNotFound)
}

// Now say we wanna take some inputs from the user eg the template, parsingLink function, username, password. One way is to make a struct of these and accept the struct. User may pass, may not.
// downside is we may conditions that we need all 3 eg connString, password, username.
// We also may have certain scenario say if one is true other can't me false kind of dependency.

// solution is functional options
// https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis

type HandlerOption func(h *handler)

func WithTemplates(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.t = t
	}
}
func WithDBConnection(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.t = t
	}
}

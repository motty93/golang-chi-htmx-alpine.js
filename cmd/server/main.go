package main

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var tmpls *template.Template

func main() {
	tmpls = template.Must(template.ParseGlob("templates/*.html.tmpl"))
	tmpls = template.Must(tmpls.ParseGlob("templates/partials/*.html.tmpl"))

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	r.Get("/hello", HelloHandler)
	r.Get("/", IndexHandler)
	r.Get("/counter", CounterHandler)

	http.ListenAndServe(":8888", r)
}

// TODO: handlerはinternal/handlersへ
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]int{"Count": 1}

	tmpls.ExecuteTemplate(w, "index.html.tmpl", data)
}

func CounterHandler(w http.ResponseWriter, r *http.Request) {
	count := r.URL.Query().Get("count")
	if count == "" {
		count = "0"
	}

	c1, err := strconv.Atoi(count)
	if err != nil {
		http.Error(w, "Invalid count parameter", http.StatusBadRequest)
		return
	}

	data := map[string]int{"Count": c1}

	tmpls.ExecuteTemplate(w, "counter.html.tmpl", data)
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Article Struct
type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func main() {
	type Articles []Article
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		articles := Articles{
			Article{Title: "Tired", Desc: "❤️", Content: "😩"},
			Article{Title: "Hey", Desc: "🤔", Content: "🙂"},
		}

		json.NewEncoder(w).Encode(articles)
	})

	r.HandleFunc("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]

		fmt.Fprintf(w, "Hello %s\n", name)
	})

	http.ListenAndServe(":9990", r)
}

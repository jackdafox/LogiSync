package http

import (
	"fmt"
	"logisync/http/db"
	"net/http"

	"github.com/gorilla/mux"
)

func HelloWorld() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":80", nil)

}

func DynamicRequest() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", nil)
}

func Router() {
	r := mux.NewRouter()
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			title := vars["title"] // the book title slug
			page := vars["page"]  // the page

			fmt.Fprintf(w, "You've requested book on: %s on page %s\n", title, page)
		})

	http.ListenAndServe(":80", r)

}

func CreateDB() {
	db.InitDB()
}

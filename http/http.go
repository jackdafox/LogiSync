package httpfunc

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"logisync/http/db"
	"net/http"
)

type Person struct {
	User struct {
		Name     string `json:"name"`
		SiteRole string `json:"siteRole"`
	} `json:"user"`
}

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
		page := vars["page"]   // the page

		fmt.Fprintf(w, "You've requested book on: %s on page %s\n", title, page)
	})

	http.ListenAndServe(":80", r)

}

func BodyTest() {
	r := mux.NewRouter()
	r.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal("Error reading response body:", err)
		}
		var person Person
		errJson := json.Unmarshal(body, &person)
		if errJson != nil {
			log.Fatal(err)
		}
		fmt.Println(person.User.SiteRole)

	})

	http.ListenAndServe(":80", r)
}

func CreateDB() {
	db.InitDB()
}

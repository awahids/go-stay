
// main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Article - Our struct for all articles
type Article struct {
	Id      string    `json:"Id"`
	Title   string `json:"Title"`
}

type Category struct {
	Id      string    `json:"Id"`
	Name string `json:"Name"`
}

var Articles []Article
var Categories []Category

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnAllCategories(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllCategories")
	json.NewEncoder(w).Encode(Categories)
}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/articles", returnAllArticles)
		myRouter.HandleFunc("/categories", returnAllCategories)
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Articles = []Article{
		Article{Id: "1", Title: "Hello"},
		Article{Id: "2", Title: "Hello 2"},
	}
	Categories = []Category{
		Category{Id: "1", Name: "succes-story"},
	}
	handleRequests()
}
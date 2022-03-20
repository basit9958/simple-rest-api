package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Song struct {
	Title   string `json:"Title"`
	Desc    string `json:"dec"`
	Content string `json:"content"`
}

type Songs []Song

func allSongs(w http.ResponseWriter, r *http.Request) {
	articles := Songs{
		Song{Title: "Test Title", Desc: "Test Description", Content: "Test Content"},
	}
	fmt.Println("Endpoint Hit: All Songs Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func testpostsongs(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Post Endpoint Hit")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	myrouter := mux.NewRouter().StrictSlash(true)
	myrouter.HandleFunc("/", homePage)
	myrouter.HandleFunc("/songs", allSongs).Methods("GET")
	myrouter.HandleFunc("/songs", testpostsongs).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myrouter))
}
func main() {
	handleRequests()
}

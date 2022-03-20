package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/article",allSongs)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
func main() {
	handleRequests()
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const PORT = ":10000"

type Article struct {
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func main() {

	Articles = []Article{
		{Title: "Hello", Desc: "seongwon", Content: "name"},
		{Title: "World", Desc: "lucas", Content: "nickname"},
	}

	log.Println("server running at", PORT)
	handleRequest()
}

func homepage(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w, "Welcome to the seongwon")
	fmt.Println("Endpoint Hit: homepage")

}

func returnAllArticle(w http.ResponseWriter, r *http.Request)  {

	fmt.Println("Endpoint Hit : return all article")
	json.NewEncoder(w).Encode(Articles)
}

func handleRequest()  {

	http.HandleFunc("/", homepage)
	http.HandleFunc("/articles", returnAllArticle)
	log.Fatal(http.ListenAndServe(PORT, nil))

}

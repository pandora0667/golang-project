package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = ":10000"

func main() {

	log.Println("server running at", PORT)
	handleRequest()
}

func homePage(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w, "Welcome to the seongwon")
	fmt.Println("Endpoint Hit: homePage")

}

func handleRequest()  {

	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(PORT, nil))

}

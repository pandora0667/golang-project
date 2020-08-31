package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

const PORT = ":8081"

func main()  {

	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {

		fmt.Fprintf(writer, "Hello %q", html.EscapeString(request.URL.Path))

	})

	http.HandleFunc("/name", func(writer http.ResponseWriter, request *http.Request) {

		fmt.Fprintf(writer, "seongwon")

	})

	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Println("Server Running at " + PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

// example http error code
func PageNotFound(w http.ResponseWriter, r *http.Response) {

	http.Error(w, "page not found try again", http.StatusNotFound)
	w.WriteHeader(http.StatusNotFound)

}

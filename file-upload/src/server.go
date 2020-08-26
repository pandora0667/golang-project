package src

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const PORT = ":8080"

func Init()  {

	http.HandleFunc("/upload", UploadFile)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Println("The server failed to run. Check the following log.")
		log.Fatal(err)
	}

	http.ListenAndServe(PORT, nil)

}

func UploadFile(w http.ResponseWriter, r *http.Request)  {

	fmt.Println("file upload endpoint hit")

	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		log.Println("Error Retrieving the File ")
		log.Println(err)
		return
	}

	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
	
	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		log.Println(err)
	}
	
	defer tempFile.Close()
	
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
	}
	
	defer tempFile.Close()

	tempFile.Write(fileBytes)
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}


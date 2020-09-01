package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


type User struct {
	gorm.Model
	Name string
	Email string
}

func main()  {

	log.Println("golang orm start")

	InitMigration()
	HandleRequest()

}

func AllUsers (w http.ResponseWriter, r *http.Request)  {
	
	//fmt.Fprintf(w, "All User")

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		panic("failed to connect database")
	}

	var users []User
	db.Find(&users)
	fmt.Println("{}", users)

	json.NewEncoder(w).Encode(users)
	
}

func NewUsers(w http.ResponseWriter, r *http.Request)  {
	
	//fmt.Fprintf(w, "User Update")

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		panic("failed to connect database")
	}

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{Name: name, Email: email})
	fmt.Fprintf(w, "user infomation create complete")
	
}

func UpdateUser(w http.ResponseWriter, r *http.Request)  {

	//fmt.Fprintf(w, "Update User")

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		panic("failed to connect database")
	}

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User
	db.Where("name = ?", name).Find(&user)
	user.Email = email
	db.Save(&user)

	fmt.Fprintf(w, "user email update complete")


}

func DeleteUser(w http.ResponseWriter, r *http.Request)  {

	//fmt.Fprintf(w, "Delete User")

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		panic("failed to connect database")
	}

	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, "user information delete complete")

}

func HandleRequest()  {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/users", AllUsers).Methods("GET")
	router.HandleFunc("/user/{name}/{email}", NewUsers).Methods("POST")
	router.HandleFunc("/user/{name}/{email}", UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{name}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

}

func InitMigration() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})
	log.Println("sqlite database connect")

}




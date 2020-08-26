package main

import (
	. "./src"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
)

func main()  {

	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != err {
		log.Println(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	fmt.Println(reflect.TypeOf(responseData))

	if err != err {
		log.Fatal(err) //에러 로그를 출력하고 프로그램을 완전히 종료
	}

	fmt.Println(string(responseData))

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Name)
	fmt.Println(len(responseObject.Pokemon))

	for i := 0; i < len(responseObject.Pokemon); i++ {
		fmt.Println(responseObject.Pokemon[i].Species.Name)
	}
}

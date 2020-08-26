package main

import (
	f "fmt"
)

func main() {

	f.Println("지연 호출")
	name()

	array := []int{1, 2, 3, 4, 5}

	for _, value := range array {
		defer f.Println(value)
	}

}

func name() {
	defer func() {
		f.Println("lee")
	}()

	func() {
		f.Println("seongwon")
	}()
}


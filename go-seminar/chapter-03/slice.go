package main

import f "fmt"

func main() {

	originalArray := [...]int{1, 2, 3}
	var cloneArray [3]int

	cloneArray = originalArray
	cloneArray[0] = 4
	f.Println("배열 복사 : ", "original : ", originalArray, "clone : ", cloneArray)

	originalSlice := []int{1, 2, 3}
	var cloneSlice []int

	cloneSlice = originalSlice
	cloneSlice[0] = 4
	f.Println("슬라이스 복사 : ", "original : ", originalSlice, "clone : ", cloneSlice)

	f.Println("-------------------------")

	a := []string{"windows", "linux", "android", "ios"}
	b := make([]string, 3)

	copy(b, a)
	f.Println(a)
	f.Println(b)

	b[0] = "ubuntu"
	f.Println(a)
	f.Println(b)

}


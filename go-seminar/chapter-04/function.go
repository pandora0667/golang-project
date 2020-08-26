package main

import f "fmt"

func main() {

	f.Println("가변 함수")

	result := multiplication(2, 3, 4, 5)
	f.Println(result)

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f.Println(sliceSum(slice...))

	r := func(a int, b int) int {
		return a + b
	}(1, 3)
	f.Println(r)

}

func multiplication(n ...int) int {

	calculation := 2
	for _, value := range n {
		calculation *= value
	}

	return calculation
}

func sliceSum(n ...int) int {

	result := 0
	for _, value := range n {
		result += value
	}

	return result
}


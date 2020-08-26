package main

import (
	f "fmt"
)

func main()  {

	var number int
	f.Println("Multiplication table")
	f.Scanln(&number)

	for i := 0; i < 10; i++ {
		f.Println(number, "*", i, "=", number * i)
	}

	f.Println("--------------------")

	for i := 0; i < 5; i++ {
		if i == 3 {
			f.Println("3이 되었습니다.")
			break
		}
	}

	f.Println("--------------------")

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				continue
			}
			f.Println(i, j)
		}

	}
}

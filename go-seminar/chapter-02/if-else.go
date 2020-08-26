package main

import f "fmt"

func main()  {

	var num int
	f.Println("1~10 정수 입력")
	f.Scanln(&num)

	if num >= 1 && num <=5 {
		f.Println("숫자는 1보다 크고 5보다 작다")
	} else if num >= 6 && num <=10 {
		f.Println("숫자는 6보다 크고 10보다 작다")
	} else {
		f.Println(num)
	}
}

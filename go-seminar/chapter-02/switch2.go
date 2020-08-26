package main

import f "fmt"

func main()  {

	var number int
	f.Println("1~10 정수 입력")
	f.Scanln(&number)

	if number <= 10{
		switch number {
		case 2, 4, 6, 8, 10:
			f.Println("짝수")
		default:
			f.Println("홀수")
		}
	} else {
		f.Println("잘못된 입력")
	}

}

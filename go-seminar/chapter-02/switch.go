package main

import f "fmt"

func main()  {

	var number int
	f.Println("Input Number")
	f.Scanln(&number)

	switch number {
	case 1:
		f.Println("1")
	case 2:
		f.Println("2")
	case 3:
		f.Println("3")
	default:
		f.Println(number)
	}

	word := "go"

	switch word {
	case "go":
		f.Println("go")
	case "world":
		f.Println("world")
	default:
		f.Println("not match")
	}

	test := 1

	switch test {
	case 1:
		f.Println(1)
		fallthrough
	case 2:
		f.Println(2)
		fallthrough
	default:
		f.Println("default")
	}
}

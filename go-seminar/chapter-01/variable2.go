package main

import f "fmt"

type ByteSize uint64

var name, title, num1, num2 = "seongwon", "golang", 1, 2

const NICKNAME = "lucas"

const (
	GO = iota
	JAVA
	PYTHON
	C
	CPP
)

const (
	_ = iota
	KB ByteSize = 1 << (10 * iota) // 1024
	MB
	GB
	TB
	PT
	EB
)

func main()  {
	i, b, s := 1, 3.141592, "example"
	f.Println(i, b, s)

	f.Println("nickname : ", NICKNAME)
	f.Println(name, title, num1, num2)

	f.Println(GO, JAVA, PYTHON, C, CPP)

	f.Println(KB, MB, GB, TB, PT, EB)
}



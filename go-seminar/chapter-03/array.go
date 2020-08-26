package main

import f "fmt"

func main() {

	original := [...]string{"ios", "linux", "unix", "android", "windows"}
	clone := original // 배열 전체 복사

	original[3] = "ubuntu"

	f.Println(original)
	f.Println(clone)
}
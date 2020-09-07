package main

import (
	f "fmt"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	f.Println("나의 코어수 : ", runtime.NumCPU())

	f.Println("goroutine 경량 실행 스레드")

	go Hello()


	//time.Sleep(time.Second)
	f.Println("seongwon")

	for i := 0; i <3; i++ {
		go func(n int) {
			f.Println("고루틴", n)
		}(i)
	}
	f.Scanln()
}

func Hello() {

	f.Println("Hello World")

}


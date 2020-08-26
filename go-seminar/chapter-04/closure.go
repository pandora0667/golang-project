package main
import f "fmt"

func main() {

	nextInt := intSeq()

	f.Println(nextInt())
	f.Println(nextInt())
	f.Println(nextInt())

	newInts := intSeq()
	f.Println(newInts())

	nextInt = intSeq()
	f.Println(nextInt())
}

func intSeq() func() int {

	i := 0
	return func() int {
		i += 1
		return i
	}

}

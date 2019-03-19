package base

import "fmt"

func intSeq() func() int {
	i := 0
	fmt.Println("aaa",i)
	return func() int {
		i++
		return i
	}
}
func main() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}

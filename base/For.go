package base

import "fmt"

func main(){
	i:=1
	for i<3{
		fmt.Println(i)
		i++
	}
	for j:=7;j<9;j++{
		fmt.Println(j)
	}
	for {
		fmt.Println("loop")
		break
	}

	// You can also `continue` to the next iteration of
	// the loop.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
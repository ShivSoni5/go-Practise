package main

import "fmt"

func zeroval(ival string) {
	ival = "fuck you"
}

func zeroptr(iptr *string) {
	*iptr = "fuck you"
}

func main() {
	i := "fuck off"
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}

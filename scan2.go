package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
//	fmt.Println(scanner.Scan()) // scanner.Scan().Text() -> error
}

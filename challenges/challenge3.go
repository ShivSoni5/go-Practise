package main

import "fmt"

func main() {
	var input int
	fmt.Scan(&input)

	if input%2 == 0 {
		output := InBetween(input)
		fmt.Println(output)
	} else {
		fmt.Println("Weird")
	}
}

func InBetween(num int) string {
	if (num >= 2) && (num <= 5) {
		return "Not Weird"
	} else if (num >= 6) && (num <= 20) {
		return "Weird"
	} else {
		return "Not Weird"
	}
}


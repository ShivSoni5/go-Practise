package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)   // first method
    fmt.Println("Simple Shell")
    fmt.Println("---------------")
    text,_ := reader.ReadString('\n')
    fmt.Println(text)

    text = strings.Replace(text, "\n", "", -1)

    if strings.Compare("hi", text) == 0 {
	    fmt.Println("hello, Yourself")
    }
}


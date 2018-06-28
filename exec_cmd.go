package main

import (
	"fmt"
	"bufio"
	"os/exec"
	"os"
	"runtime"
)

func execute(cmd string) {
	out,err := exec.Command(cmd).Output()

	if err != nil {
		fmt.Printf("%s \n",err)
	}
	fmt.Println("Command Successfully Executed")

	output := string(out[:])
	fmt.Println(output)
}

func main() {
	fmt.Println("Simple Shell")
	fmt.Println("-------------")
	if runtime.GOOS == "windows" {
		fmt.Println("Can't execute linux commands here")
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		execute(scanner.Text())
	}
}

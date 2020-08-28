package main

import "fmt"

func main() {
	fmt.Printf("%s\n", "Hello World")
	fmt.Printf("%s\n", helloWorld(""))
}

func helloWorld(s string) string {
	if s == "lower" {
		return "hello world"
	} else {
		return "HELLO WORLD"
	}
}

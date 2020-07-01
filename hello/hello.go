package main

import "fmt"

// Hello returns our desired message
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text int) []string {
	return make([]string, text)
}

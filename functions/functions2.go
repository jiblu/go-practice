package main

import "fmt"

func and(a, b string) string {
	return a + b
}

func main() {
	string := and("testing", " 123")
	fmt.Println(string)
}


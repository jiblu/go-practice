package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"

	fmt.Println("set:", s)
	s[1] = "b"
	fmt.Println("set:", s)
	s[2] = "c"
	fmt.Println("set:", s)

	s = append(s, "d")
	fmt.Println("apd:", s)
	
}
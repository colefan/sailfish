package main

import "fmt"

func main() {
	xxx := make([]byte, 2, 2)
	fmt.Printf("xxx len = %d,cap = %d \n", len(xxx), cap(xxx))
	xxx = append(xxx[0:0], 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(xxx)
}

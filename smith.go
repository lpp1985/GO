package main

import (
	"fmt"
	"lpp"
)

func main() {
	A := "ATCGGGCCAA"
	B := "ATCGGGCCAT"
	C := "ATCGGGCCAA"
	fmt.Println(lpp.SmithWaterman(A, B))
	fmt.Println(lpp.SmithWaterman(A, C))
}

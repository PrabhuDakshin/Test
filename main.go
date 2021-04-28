package main

import (
	ad "github.com/PrabhuDakshin/Test/add"
	"os"
	"strconv"
)

func main() {
	a := os.Args[1]
	b := os.Args[2]
	A, _ := strconv.Atoi(a)
	B, _ := strconv.Atoi(b)
	ad.ADD(A, B)
}


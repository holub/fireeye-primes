package main

import (
	"os"
	"strconv"
	"fmt"
	"github.com/holub/fireeye-primes/primes"
	"github.com/holub/fireeye-primes/calculator"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("usage: %s <ubound>\n", os.Args[0])
		os.Exit(2)
	}

	ubound, err := strconv.Atoi(os.Args[1:][0])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	if ubound < 1 {
		fmt.Println("Please use upper bound > 0")
		os.Exit(1)
	}

	p := primes.New(ubound)
	calculator.DoCalculateRanges(p)

	os.Exit(0)
}

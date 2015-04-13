package calculator

import (
	"fmt"
	"github.com/holub/fireeye-primes/primes"
)

// Read ranges and print statistics for each
func DoCalculateRanges(pr *primes.Primes) {
	var err error
	var lb, ub int
	for {
		lb, err = readUint("Enter a lower bound: ")
		if err == nil {
			ub, err = readUint("Enter a upper bound: ")
			if err == nil {
				err = printCalc(pr, lb, ub)
			}
		}

		// Any errors?
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}
}

// Read and check if input is positive number
func readUint(prompt string) (int, error) {
	fmt.Print(prompt)

	var i int
	_, err := fmt.Scan(&i)
	if err != nil {
		i = -1
	} else if i < 1 {
		err = fmt.Errorf("expected positive number")
		i = -1
	}

	return i, err;
}

// TODO: split Calc and print to make Calc testable
// Print statistics for range
func printCalc(pr *primes.Primes, lb int, ub int) error {
	var err error
	if ub > pr.Ubound {
		err = fmt.Errorf("ERR upper bound > %d", pr.Ubound)
	} else if lb > ub {
		err = fmt.Errorf("ERR lower > upper")
	} else {
		var sum int64 = 0
		var count int = 0
		var cur int
		fmt.Println("Result:")
		fmt.Print("Prime numbers: [")
		for n := lb; n > 0 && n <= ub; n = cur + 1 {
			cur = pr.Store.FindFrom(n)
			if (cur > 0 && cur <= ub) {
				if sum > 0 {
					fmt.Print(", ")
				}
				fmt.Print(cur)
				sum = sum + int64(cur)
				count++
			}
		}
		fmt.Println("]")
		fmt.Println("Sum: ", sum)
		fmt.Printf("Mean: %.2f\n", float64(sum) / float64(count)) 
	}

	return err
}

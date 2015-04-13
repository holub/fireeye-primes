package primes

import (
	"fmt"
	"math"
	"log"
	"github.com/holub/fireeye-primes/redisstore"
)

func New(ubound int) *Primes {
	pr := &Primes{
		Ubound: ubound,
		Store: redisstore.New(),
	}
	pr.generate()

	return pr
}

// Implementation of Sieve of Atkin - https://en.wikipedia.org/wiki/Sieve_of_Atkin
// ...and http://stackoverflow.com/questions/10580159/sieve-of-atkin-explanation-and-java-example
func (pr *Primes) generate() {
	log.Printf("Generating prime numbers between 1 and %d", pr.Ubound)

	limit := pr.Ubound
	limitSqrt := int(math.Sqrt(float64(limit)))

	pr.Store.Reset()
	//pr.Store.Add(1)
	pr.Store.Add(2)
	pr.Store.Add(3)


    for x := 1; x <= limitSqrt; x++ {
        for y := 1; y <= limitSqrt; y++ {
            // first quadratic using m = 12 and r in R1 = {r : 1, 5}
            n := (4 * x * x) + (y * y)
            if n <= limit && (n % 12 == 1 || n % 12 == 5) {
                pr.Store.Invert(n)
            }
            // second quadratic using m = 12 and r in R2 = {r : 7}
            n = (3 * x * x) + (y * y)
            if n <= limit && (n % 12 == 7) {
                pr.Store.Invert(n)
            }
            // third quadratic using m = 12 and r in R3 = {r : 11}
            n = (3 * x * x) - (y * y)
            if x > y && n <= limit && (n % 12 == 11) {
                pr.Store.Invert(n)
            }
            // note that R1 union R2 union R3 is the set R
            // R = {r : 1, 5, 7, 11}
            // which is all values 0 < r < 12 where r is 
            // a relative prime of 12
            // Thus all primes become candidates
        }
    }
    // remove all perfect squares since the quadratic
    // wheel factorization filter removes only some of them
    for n := 5; n <= limitSqrt; n++ {
        if pr.Store.Get(n) == 1 {
            x := n * n
            for i := x; i <= limit; i += x {
                pr.Store.Set(i, 0)
            }
        }
    }
}

// TODO: remove
func (pr *Primes) PrintPretty() {
	limit := pr.Ubound
	j := 0
    for i := 0; i <= limit; i++ {
        if pr.Store.Get(i) == 1 {
			fmt.Printf("%8d ", i)
			j = j + 1
            if j % 10 == 0 {
				fmt.Println()
            }
            if j % 100 == 0 {
				fmt.Println()
            }
        }
    }
	fmt.Println()
}

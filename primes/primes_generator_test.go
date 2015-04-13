package primes

import "testing"


func TestPrimes(t *testing.T) {

	p := New(101)
	s := p.GetStore()

	primes := []int{2, 3, 5, 7, 97, 101}
	for _, i := range primes {
		if 1 != s.Get(i) {
			t.Error("Prime number not found at:", i)
		}
	}

	notPrimes := []int{0, 1, 4, 6, 100, 102}
	for _, i := range notPrimes {
		if 0 != s.Get(i) {
			t.Error("Unexpected prime at:", i)
		}
	}

}

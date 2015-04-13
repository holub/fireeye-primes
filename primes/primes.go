package primes

import "github.com/holub/fireeye-primes/redisstore"


type Primes struct {
	Ubound int
	Store *redisstore.Store
}

# Description
Create a NodeJS, Scala, or Golang (pick one) application that accomplishes the following:

- It should generate all the prime numbers between 1 to X where X is a command line argument to the application

Example for NodeJS: node app.js 100

- The prime numbers should be stored in a local Redis instance
- Once the prime numbers are generated the application should repeatedly ask the user for a lower and upper bounds (inclusive) on the prime numbers to return along with their sum and mean

## Example flow:

	$ Enter a lower bound: 3
	$ Enter an upper bound: 9
	$ Result:
	$ Prime numbers: [3, 5, 7]
	$ Sum: 15
	$ Mean: 5


- Include basic unit testing around the core functionality
- All code (excluding any external dependencies) should be committed to a GitHub repository.

# Requirements
- [Redis](http://redis.io/): 2.8.19
- [Go](https://golang.org/): 1.4.2

# Execution

	$ go run fireeyeprimes.go 1000

# Implementation

[Sieve of Atkin](http://en.wikipedia.org/wiki/Sieve_of_Atkin)

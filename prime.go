package main

import (
	"encoding/gob"
	"errors"
	"os"
)

/**
	This project uses Sieve of Sundaram which was discovered by an Indian mathematician Mr.S.P Sundaram in 1934
	for more information about the core algorithm, refer the following url:
	https://en.wikipedia.org/wiki/Sieve_of_Sundaram
 */

// define Prime as a slice of integer
type Prime []int

func NewPrime() Prime {
	return make(Prime, 0)
}

// SieveOfSundaram receives n as an integer number. It will base on Sundaram algorithm get all primes in n's range.
func (p *Prime)SieveOfSundaram(n int) {
	if n < 2 {
		return
	}
	// In general, Sieve of Sundaram produces less than (2x+2) number of primes.
	// <=> 2x+2 <= n
	// <=> x <= (n-2)/2
	k := (n-2)/2
	a := make([]bool, k+1)
	*p = append(*p, 2)

	for i:=1; i<k+1; i++{
		for j:=i; i+j+(2*i*j) <= k; j++ {
			a[i+j+(2*i*j)] = true
		}
	}
	for i:=1; i<k+1; i++ {
		if !a[i] {
			*p = append(*p, 2*i+1)
		}
	}
}

// Save saves list primes into file
func (p Prime)SavePrimes(fileName string) error {
	if len(p) == 0 {
		return errors.New("primes is empty")
	}
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	return encoder.Encode(p)
}

// GetPrimes reads primes from given fileName
func (p Prime)GetPrimes(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	return decoder.Decode(&p)
}

package prime

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
type Prime []int32

func NewPrime() Prime {
	return make(Prime, 0)
}

// SieveOfSundaram receives n as an integer number. It will base on Sundaram algorithm get all primes in n's range.
func SieveOfSundaram(n int) Prime {
	if n < 2 {
		return nil
	}
	// In general, Sieve of Sundaram produces less than (2x+2) number of primes.
	// <=> 2x+2 <= n
	// <=> x <= (n-2)/2
	k := (n-2)/2
	a := make([]bool, k+1)

	// init primes which also has k+1 length
	primes := make([]int32, k+1)
	primes[0] = 2
	// count how many primes
	for i:=1; i<k+1; i++{
		for j:=i; i+j+(2*i*j) <= k; j++ {
			a[i+j+(2*i*j)] = true
		}
	}
	counter := 1
	for i:=1; i<k+1; i++ {
		if !a[i] {
			primes[counter] = int32(2*i+1)
			counter ++
		}
	}
	// get valuable parts
	return primes[0:counter]
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
func (p *Prime)GetPrimes(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	return decoder.Decode(&p)
}

// BinarySearch is used to search highest prime that is less than given number (n)
func (p Prime)BinarySearch(left, right, n int32) int32 {
	if len(p) == 0 {
		return 0
	}
	if left <= right {
		mid := (left+right)/2
		// if mid reaches left corner (0) or right corner (len-1) then return primes[mid]
		if mid == 0 || mid == int32(len(p)-1) {
			return p[mid]
		}
		// if primes[mid] is n then return primse[mid-1] which is previous element.
		if p[mid] == n {
			return p[mid-1]
		}
		if p[mid] < n && p[mid+1] > n {
			return p[mid]
		}
		if p[mid] > n {
			return p.BinarySearch(left, mid - 1, n)
		}
		return p.BinarySearch(mid + 1, right, n)
	}
	return 0
}

// LoadPrimes loads primes from file and if file does not exist, call SieveOfSundaram to get all primes.
func (p *Prime)LoadPrimes(fileName string, number int) error {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		// start SieveOfSundaram in `number's range` and save to file
		*p = SieveOfSundaram(number)
		return p.SavePrimes(fileName)
	}
	return p.GetPrimes(fileName)
}

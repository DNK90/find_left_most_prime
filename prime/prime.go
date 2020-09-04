package prime

import (
	"errors"
	"github.com/dnk90/find_left_most_prime/proto"
	proto2 "github.com/gogo/protobuf/proto"
	"io/ioutil"
	"os"
)

/**
	This project uses Sieve of Sundaram which was discovered by an Indian mathematician Mr.S.P Sundaram in 1934
	for more information about the core algorithm, refer the following url:
	https://en.wikipedia.org/wiki/Sieve_of_Sundaram
 */

// define Prime as a slice of integer
type Prime proto.Prime

func NewPrime() *Prime {
	return &Prime{Primes: make([]int32, 0)}
}

// SieveOfSundaram receives n as an integer number. It will base on Sundaram algorithm get all primes in n's range.
func SieveOfSundaram(n int) *Prime {
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
	return &Prime{Primes: primes[0:counter]}
}

// Save saves list primes into file
func (p *Prime)SavePrimes(fileName string) error {
	if len(p.Primes) == 0 {
		return errors.New("primes is empty")
	}
	//file, err := os.Create(fileName)
	//if err != nil {
	//	return err
	//}
	//defer file.Close()
	data, err := proto2.Marshal(&proto.Prime{Primes: p.Primes})
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, data, 0777)
	//encoder := proto.NewEncoder(file)
	//return encoder.Encode(p)
}

// GetPrimes reads primes from given fileName
func (p *Prime)GetPrimes(fileName string) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	var prime proto.Prime
	if err = proto2.Unmarshal(data, &prime); err != nil {
		return err
	}
	p.Primes = prime.Primes
	return nil
}

// BinarySearch is used to search highest prime that is less than given number (n)
func (p Prime)BinarySearch(left, right, n int32) int32 {
	if len(p.Primes) == 0 {
		return 0
	}
	if left <= right {
		mid := (left+right)/2
		// if mid reaches left corner (0) or right corner (len-1) then return primes[mid]
		if mid == 0 || mid == int32(len(p.Primes)-1) {
			return p.Primes[mid]
		}
		// if primes[mid] is n then return n since n is a prime.
		if p.Primes[mid] == n {
			return n
		}
		if p.Primes[mid] < n && p.Primes[mid+1] > n {
			return p.Primes[mid]
		}
		if p.Primes[mid] > n {
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
		p = SieveOfSundaram(number)
		return p.SavePrimes(fileName)
	}
	return p.GetPrimes(fileName)
}

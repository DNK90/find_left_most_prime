package main

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"math"
	"os"
	"path/filepath"
	"testing"
	"time"
)

type PrimeTestSuite struct {
	suite.Suite
	fileName string
	prime Prime
}

func (suite *PrimeTestSuite) SetupSuite() {
	dir, err := os.Getwd()
	suite.NoError(err)
	suite.fileName = filepath.Join(dir, "testFile")
	suite.prime = NewPrime()
}

func (suite *PrimeTestSuite) TearDownSuite() {
	// remove fileName
	os.Remove(suite.fileName)
}

func(suite *PrimeTestSuite)Test_1_LoadPrimes_1_LessThan2() {
	suite.prime.SieveOfSundaram(1)
	suite.Equal(0, len(suite.prime))
}

func (suite *PrimeTestSuite)Test_1_LoadPrimes_2_GreaterThan2() {
	expected := Prime([]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113})
	startTime := time.Now()
	suite.prime.SieveOfSundaram(math.MaxInt8)
	endTime := time.Now()
	println(fmt.Sprintf("Elapsed time is %v len is %v", endTime.Sub(startTime), len(suite.prime)))
	suite.Equal(expected, suite.prime)
}

func (suite *PrimeTestSuite)Test_2_SavePrimes() {
	err := suite.prime.SavePrimes(suite.fileName)
	suite.NoError(err)
}

func (suite *PrimeTestSuite)Test_2_SavePrimes_WithEmptyPrimes() {
	prime := NewPrime()
	err := prime.SavePrimes(suite.fileName)
	suite.Errorf(err, "primes is empty")
}

func (suite *PrimeTestSuite)Test_3_GetPrimesFromFile() {
	expected := Prime([]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113})
	err := suite.prime.GetPrimes(suite.fileName)
	suite.NoError(err)
	suite.Equal(expected, suite.prime)
}

func (suite *PrimeTestSuite)Test_3_GetPrimesFromFile_FileDoesNotExist() {
	err := suite.prime.GetPrimes("notExistFile")
	suite.Error(err)
}

func TestPrime(t *testing.T) {
	suite.Run(t, new(PrimeTestSuite))
}

package prime

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
	fileName, startFile string
	prime               Prime
}

func (suite *PrimeTestSuite) SetupSuite() {
	dir, err := os.Getwd()
	suite.NoError(err)
	suite.fileName = filepath.Join(dir, "testFile")
	suite.startFile = filepath.Join(dir, "startFile")
	suite.prime = NewPrime()
}

func (suite *PrimeTestSuite) TearDownSuite() {
	// remove fileName
	os.Remove(suite.fileName)
	os.Remove(suite.startFile)
}

func(suite *PrimeTestSuite)Test_1_LoadPrimes_1_LessThan2() {
	suite.prime = SieveOfSundaram(1)
	suite.Equal(0, len(suite.prime))
}

func (suite *PrimeTestSuite)Test_1_LoadPrimes_2_GreaterThan2() {
	expected := Prime([]int32{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113})
	startTime := time.Now()
	suite.prime = SieveOfSundaram(math.MaxInt8)
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
	expected := Prime([]int32{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113})
	err := suite.prime.GetPrimes(suite.fileName)
	suite.NoError(err)
	suite.Equal(expected, suite.prime)
}

func (suite *PrimeTestSuite)Test_3_GetPrimesFromFile_FileDoesNotExist() {
	err := suite.prime.GetPrimes("notExistFile")
	suite.Error(err)
}

func (suite *PrimeTestSuite)Test_4_LoadPrimes_1_FileDoesNotExist() {
	_, err := os.Stat(suite.startFile)
	suite.True(os.IsNotExist(err))
	err = suite.prime.LoadPrimes(suite.startFile, math.MaxInt16)
	suite.NoError(err)
	// startFile must be created after LoadPrimes finish
	_, err = os.Stat(suite.startFile)
	suite.False(os.IsNotExist(err))
	suite.True(len(suite.prime) > 0)
}

func (suite *PrimeTestSuite)Test_4_LoadPrimes_2_FileExist() {
	prime := NewPrime()
	_, err := os.Stat(suite.startFile)
	suite.False(os.IsNotExist(err))
	err = prime.LoadPrimes(suite.startFile, math.MaxInt16)
	suite.NoError(err)
	suite.True(len(suite.prime) > 0)
}

func (suite *PrimeTestSuite)Test_5_BinarySearch() {
	expected := int32(59)
	actual := suite.prime.BinarySearch(0, int32(len(suite.prime)-1), 60)
	suite.Equal(expected, actual)
}

func (suite *PrimeTestSuite)Test_5_BinarySearch_EmptyPrimes() {
	prime := NewPrime()
	suite.Equal(int32(0), prime.BinarySearch(0, int32(len(prime)), 1))
}

func (suite *PrimeTestSuite)Test_5_BinarySearch_LeftGreaterThanRight() {
	suite.Equal(int32(0), suite.prime.BinarySearch(1, 0, 1))
}

func TestPrime(t *testing.T) {
	suite.Run(t, new(PrimeTestSuite))
}

package main
import "testing"
import "reflect"
import "math"

type testPair struct {
  value int
  sequence []int
}
type testPairOverflowed struct {
  value1 int
  value2 int
  overflow bool
}
/**
* Test for integer overflow
**/
var testsOverflow = []testPairOverflowed {
	{ 20, 20, false},
	{ math.MaxUint32, math.MaxUint32, true},
	{ 1, math.MaxUint32, true},
	{ 0, math.MaxUint32, false},
}
func testsOverflowed(t *testing.T) {
  for _, pair := range testsOverflow {
    v := overflowed(pair.value1, pair.value2)
    if v != pair.overflow {
      t.Error(
        "For", pair.value1, pair.value2,
        "expected", pair.overflow,
        "got", v,
      )
    }
  }
}

/**
* Test to determine fibonacci algorithm is correct 
**/
var tests = []testPair {
  {0, []int{}},
  { 1, []int{0}},
  { 13, []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144}},
}

func TestFibonacciAlgorithm(t *testing.T) {
  for _, pair := range tests {
    v := FibonacciAlgorithm(pair.value)
    if !reflect.DeepEqual(v, pair.sequence) {
      t.Error(
        "For", pair.value,
        "expected", pair.sequence,
        "got", v,
      )
    }
  }
}

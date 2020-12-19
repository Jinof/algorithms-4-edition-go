package random

import (
	"errors"
	"math/rand"
	"time"
)

var ErrorNumNegetive = errors.New("negiteve num is not allowed")

func newNums(digits, length int) []int {
	nums := make([]int, length)
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)
	for i := 0; i < length; i++ {
		nums[i] = r.Intn(digits)
	}
	return nums
}

// Generate an array with specified digits and length.
func NewNums(digits, length int) (nums []int, err error) {
	if digits < 0 || length < 0 {
		return []int{}, ErrorNumNegetive
	}
	return newNums(digits, length), nil
}

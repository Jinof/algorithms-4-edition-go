package msort

import (
	"github.com/jinof/algorithms-4-edition-go/pkg/random"
	"log"
)

var NUMS []int

const (
	DIGITS = 10
	LENGTH = 100000000
)

func init() {
	initNums()
}

func initNums() {
	var err error
	if NUMS, err = random.NewNums(DIGITS, LENGTH); err != nil && err != random.ErrorNumNegetive {
		log.Fatal("[init nums]: ", err)
	}
}

func sorted(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}

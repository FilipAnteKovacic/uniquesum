package uniquesums

import (
	"errors"
	"fmt"
)

// IntUniqueSumCheck calculate unique sum of given slice of numbers
// and check if any sum is same as given value
func IntUniqueSumCheck(nums []int64, want int64) [][]int64 {

	var combinations [][]int64

	fComs := IntCombinations(nums)

	if len(fComs) != 0 {

		for _, fCom := range fComs {

			sum, err := IntSumSlice(fCom)

			if err != nil {
				fmt.Printf("error while sum slice: %v", err)
			}

			if sum == want {
				combinations = append(combinations, fCom)
			}

		}

	}

	return combinations
}

// IntCombinations returns all combinations for a given int64 slice.
// used from https://github.com/mxschmitt/golang-combinations
func IntCombinations(set []int64) (subsets [][]int64) {
	length := uint(len(set))

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		var subset []int64

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset = append(subset, set[object])
			}
		}
		// add subset to subsets
		subsets = append(subsets, subset)
	}
	return subsets
}

// IntSumSlice return slice numbers sum
func IntSumSlice(x []int64) (int64, error) {

	if len(x) == 0 {
		return 0, errors.New("slice is empty")
	}

	var total int64
	for _, valuex := range x {
		total += valuex
	}

	return total, nil
}

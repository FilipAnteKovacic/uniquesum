package uniquesums

import (
	"errors"
	"fmt"
)

// StringUniqueSumCheck calculate unique sum of given slice of numbers
// and check if any sum is same as given value
func StringUniqueSumCheck(nums []string, want string) [][]string {

	var combinations [][]string

	fComs := StringCombinations(nums)

	if len(fComs) != 0 {

		for _, fCom := range fComs {

			sum, err := StringSumSlice(fCom)

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

// StringCombinations returns all combinations for a given string slice.
// used from https://github.com/mxschmitt/golang-combinations
func StringCombinations(set []string) (subsets [][]string) {
	length := uint(len(set))

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		var subset []string

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

// StringSumSlice return slice numbers sum
func StringSumSlice(x []string) (string, error) {

	if len(x) == 0 {
		return "", errors.New("slice is empty")
	}

	var total string
	for _, valuex := range x {
		total += valuex
	}

	return total, nil
}

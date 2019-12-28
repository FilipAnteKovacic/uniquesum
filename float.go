package uniquesums

import (
	"errors"
	"fmt"
)

// FloatUniqueSumCheck calculate unique sum of given slice of numbers
// and check if any sum is same as given value
func FloatUniqueSumCheck(nums []float64, want float64) [][]float64 {

	var combinations [][]float64

	fComs := FloatCombinations(nums)

	if len(fComs) != 0 {

		for _, fCom := range fComs {

			sum, err := FloatSumSlice(fCom)

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

// FloatCombinations returns all combinations for a given float64 slice.
// used from https://github.com/mxschmitt/golang-combinations
func FloatCombinations(set []float64) (subsets [][]float64) {
	length := uint(len(set))

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		var subset []float64

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

// FloatSumSlice return slice numbers sum
func FloatSumSlice(x []float64) (float64, error) {

	if len(x) == 0 {
		return 0, errors.New("slice is empty")
	}

	total := 0.0
	for _, valuex := range x {
		total += valuex
	}

	return total, nil
}

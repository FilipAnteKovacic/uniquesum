package uniquesums

import (
	"reflect"
	"testing"
)

func TestFloatSumSlice(t *testing.T) {

	tt := []struct {
		name  string
		slice []float64
		value float64
	}{
		{
			name:  "empty",
			slice: []float64{},
			value: 0,
		},
		{
			name:  "1+2",
			slice: []float64{1, 2},
			value: 3,
		},
		{
			name:  "1+2+3+4+5",
			slice: []float64{1, 2, 3, 4, 5},
			value: 15,
		},
		{
			name:  "10+20+30+40+50",
			slice: []float64{10, 20, 30, 40, 50},
			value: 150,
		},
		{
			name:  "11+22+33+44+55",
			slice: []float64{11, 22, 33, 44, 55},
			value: 165,
		},
	}

	for _, tc := range tt {

		t.Run(tc.name, func(t *testing.T) {

			sum, err := FloatSumSlice(tc.slice)

			if (len(tc.slice) == 0) && err == nil {
				t.Error("given empty slice not having error")
			}

			if sum != tc.value {
				t.Errorf("sum of slice expecting %v, got %v", tc.value, sum)
			}

		})

	}

}

func TestFloatCombinations(t *testing.T) {

	tt := []struct {
		name string
		in   []float64
		out  [][]float64
	}{
		{
			name: "empty",
			in:   []float64{},
			out:  nil,
		},
		{
			name: "1",
			in:   []float64{10},
			out: [][]float64{
				{10},
			},
		},
		{
			name: "2",
			in:   []float64{10, 20},
			out: [][]float64{
				{10},
				{20},
				{10, 20},
			},
		},
		{
			name: "3",
			in:   []float64{10, 20, 30},
			out: [][]float64{
				{10},
				{20},
				{10, 20},
				{30},
				{10, 30},
				{20, 30},
				{10, 20, 30},
			},
		},
		{
			name: "4",
			in:   []float64{10, 20, 30, 40},
			out: [][]float64{
				{10},
				{20},
				{10, 20},
				{30},
				{10, 30},
				{20, 30},
				{10, 20, 30},
				{40},
				{10, 40},
				{20, 40},
				{10, 20, 40},
				{30, 40},
				{10, 30, 40},
				{20, 30, 40},
				{10, 20, 30, 40},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			out := FloatCombinations(tc.in)
			if !reflect.DeepEqual(out, tc.out) {
				t.Errorf("error: \nreturn:\t%v\nwant:\t%v", out, tc.out)
			}
		})
	}
}

func TestFloatUniqueSumCheck(t *testing.T) {

	tt := []struct {
		name string
		nums []float64
		want float64
	}{
		{
			name: "2 nums",
			nums: []float64{1451, 2342},
			want: 1451,
		},
		{
			name: "5 nums",
			nums: []float64{11, 22, 33, 44, 55},
			want: 77,
		},
		{
			name: "5 nums - total sum",
			nums: []float64{11, 22, 33, 44, 55},
			want: 165,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			uSums := FloatUniqueSumCheck(tc.nums, tc.want)

			if len(uSums) == 0 {
				t.Fatal("expecting at least one unique sum")
			}

			for _, s := range uSums {

				sum, err := FloatSumSlice(s)

				if err != nil {
					t.Errorf("slice sum err: %v", err)
				}

				if sum != tc.want {
					t.Errorf("sum of slice %v expecting %v, got %v", s, tc.want, sum)
				}

			}

		})

	}

}

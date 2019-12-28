package uniquesums

import (
	"reflect"
	"testing"
)

func TestStringSumSlice(t *testing.T) {

	tt := []struct {
		name  string
		slice []string
		value string
	}{
		{
			name:  "empty",
			slice: []string{},
			value: "",
		},
		{
			name:  "FO",
			slice: []string{"F", "O"},
			value: "FO",
		},
		{
			name:  "FOOBAR",
			slice: []string{"FO", "O", "BA", "R"},
			value: "FOOBAR",
		},
		{
			name:  "HELLO",
			slice: []string{"H", "E", "L", "L", "O"},
			value: "HELLO",
		},
	}

	for _, tc := range tt {

		t.Run(tc.name, func(t *testing.T) {

			sum, err := StringSumSlice(tc.slice)

			if (len(tc.slice) == 0) && err == nil {
				t.Error("given empty slice not having error")
			}

			if sum != tc.value {
				t.Errorf("sum of slice expecting %v, got %v", tc.value, sum)
			}

		})

	}

}

func TestStringCombinations(t *testing.T) {

	tt := []struct {
		name string
		in   []string
		out  [][]string
	}{
		{
			name: "empty",
			in:   []string{},
			out:  nil,
		},
		{
			name: "1",
			in:   []string{"A"},
			out: [][]string{
				{"A"},
			},
		},
		{
			name: "2",
			in:   []string{"A", "B"},
			out: [][]string{
				{"A"},
				{"B"},
				{"A", "B"},
			},
		},
		{
			name: "3",
			in:   []string{"A", "B", "C"},
			out: [][]string{
				{"A"},
				{"B"},
				{"A", "B"},
				{"C"},
				{"A", "C"},
				{"B", "C"},
				{"A", "B", "C"},
			},
		},
		{
			name: "4",
			in:   []string{"A", "B", "C", "D"},
			out: [][]string{
				{"A"},
				{"B"},
				{"A", "B"},
				{"C"},
				{"A", "C"},
				{"B", "C"},
				{"A", "B", "C"},
				{"D"},
				{"A", "D"},
				{"B", "D"},
				{"A", "B", "D"},
				{"C", "D"},
				{"A", "C", "D"},
				{"B", "C", "D"},
				{"A", "B", "C", "D"},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			out := StringCombinations(tc.in)
			if !reflect.DeepEqual(out, tc.out) {
				t.Errorf("error: \nreturn:\t%v\nwant:\t%v", out, tc.out)
			}
		})
	}
}

func TestStringUniqueSumCheck(t *testing.T) {

	tt := []struct {
		name  string
		slice []string
		want  string
	}{
		{
			name:  "FO",
			slice: []string{"F", "O", "Q", "B", "C"},
			want:  "FO",
		},
		{
			name:  "FOOBAR",
			slice: []string{"F", "OO", "BA", "R", "BAR"},
			want:  "FOOBAR",
		},
		{
			name:  "HELLO",
			slice: []string{"H", "E", "L", "MA", "L", "O", "WO", "RL", ""},
			want:  "HELLO",
		},
		/*
			{
				name:  "ALPHBET",
				slice: []string{"A", "B", "C", "D","E", "F","G","H","I","J","K","L","M","N","O","P","R","S","T","U","V","Z","X","Y","Q"},
				want: "ALPHBET",
			},
		*/

	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			uSums := StringUniqueSumCheck(tc.slice, tc.want)

			if len(uSums) == 0 {
				t.Fatal("expecting at least one unique sum")
			}

			for _, s := range uSums {

				sum, err := StringSumSlice(s)

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

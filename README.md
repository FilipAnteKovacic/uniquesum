# UniqueSum

Set of functions that working with slices of ints, floats or strings

- Sum of given slices
- Unique combinations of given slices
- Check which combination from slice return wanted sum

## Examples

```
intSlice := []int{ 1, 2, 3, 4, 5 }
wantSum := 3

intSlices := IntUniqueSumCheck(intSlice, wantSum)

// Result:
// intSlices[0] > []int{ 1, 2 }
// intSlices[1] > []int{ 3 }

```

```
floatSlice := []float{ 2003.54, 454.32, 54.23, 653.55 }
wantSum := 2512.09

floatSlices := FloatUniqueSumCheck(floatSlice, wantSum)

// Result:
// floatSlices[0] > []int{ 2003.54, 454.32, 54.23 }

```

### Installing

```
go get github.com/FilipAnteKovacic/uniquesums
```

## Running the tests

```
go test
```

## Big thanks to

* [Max Schmitt](https://github.com/mxschmitt/golang-combinations) - using for combinations

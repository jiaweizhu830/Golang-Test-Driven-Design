package main

func Sum(numbers []int) int {
	sum := 0
	// each iteration, 'range' returns index and value
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// variadic functions (called with any number of arguments)
func SumAllByMake(numbersToSum ...[]int) []int {
	// get number of slices
	lengthOfNumbers := len(numbersToSum)
	// create a slice
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}

// variadic functions (called with any number of arguments)
func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		// append number to sums slice, and return a new slice
		sums = append(sums, Sum(numbers))
	}
	return
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			// sliced slice
			// slice [low:high]
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}

package arrays

// SumAll func
func SumAll(numbersToSum ...[]int) (sums []int) {
	lengthOfNumbers := len(numbersToSum)
	sums = make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return
}

// Sum func
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

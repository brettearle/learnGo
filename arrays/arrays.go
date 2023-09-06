package arrays

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numberSlicesToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numberSlicesToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

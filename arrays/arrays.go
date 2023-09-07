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

func SumAllTails(numsToSum ...[]int) []int {
	var sums []int
	for _, nums := range numsToSum {
		if len(nums) == 0 {
			sums = append(sums, 0)
		} else {
			tail := nums[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}

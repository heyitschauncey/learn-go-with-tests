package arraysandslices

// Sum using a fixed array
// func Sum(numbers [5]int) int {
// 	sum := 0
// 	for _, number := range numbers {
// 		sum += number
// 	}
//
// 	return sum
// }

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// this version of vars carses about the capacity of the slices
	// lengthOfNumbers := len(numbersToSum)
	// sums := make([]int, lengthOfNumbers)
	// for i, numbers := range numbersToSum {
	// sums[i] = Sum(numbers)
	// }

	// this version utilizes append(slice, newValue)
	// this will increase the capacity of the slice if necessary
	// it also starts with an empty slice, with a capacity of 0
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}

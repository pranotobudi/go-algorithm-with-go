package module01

// Sum will sum up all of the numbers passed
// in and return the result
// first way
// func Sum(numbers []int) int {
// 	var result int
// 	for _, num := range numbers {
// 		result += num
// 	}
// 	return result
// }

// second way
func Sum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	return numbers[0] + Sum(numbers[1:])
}

//

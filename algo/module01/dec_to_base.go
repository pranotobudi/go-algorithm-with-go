package module01

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//
// func DecToBase(dec, base int) string {
// 	baseRef := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}

// 	resultSlice := []string{}
// 	resultSliceReverse := []string{}
// 	dividend := dec
// 	var idx int

// 	// get result in int, in reverse order
// 	for dividend != 0 {
// 		idx = dividend % base
// 		resultSliceReverse = append(resultSliceReverse, baseRef[idx])
// 		dividend = dividend / base
// 	}

// 	// reverse the previous one
// 	for i := len(resultSliceReverse) - 1; i >= 0; i-- {
// 		resultSlice = append(resultSlice, resultSliceReverse[i])
// 	}

// 	return strings.Join(resultSlice, "")
// }

func DecToBase(dec, base int) string {
	charset := "0123456789ABCDEF"
	var res string
	for dec > 0 {
		rem := dec % base
		res = string(charset[rem]) + res
		dec = dec / base
	}
	return res
}

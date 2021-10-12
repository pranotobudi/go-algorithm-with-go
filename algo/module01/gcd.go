package module01

func GCD(a, b int) int {
	var toCheck int
	if b < a {
		toCheck = b
	} else {
		toCheck = a
	}

	var gcd int
	for i := 1; i <= toCheck; i++ {
		if a%i == 0 && b%i == 0 {
			gcd = i
		}
	}

	return gcd
}

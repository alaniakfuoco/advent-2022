package util

func Abs(i int) int {
	if i < 0 {
		return i * -1
	} else {
		return i
	}
}

func Min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func Max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

// -1 if a is less than b
//
//	0 if a equals b
//
// +1 if a is greater than b
func IntCompare(a, b int) int {
	if a < b {
		return -1
	} else if b < a {
		return 1
	} else {
		return 0
	}
}

func Power(x, y int) int {
	result := 1
	for i := 0; i < y; i++ {
		result *= x
	}
	return result
}

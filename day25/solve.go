package day25

import (
	"advent/util"
	"fmt"
	"strconv"
)

func Solve(path string) {
	data := util.GetFileLines(path)

	var sum int64
	for _, line := range data {
		// fmt.Print(line)
		// fmt.Print("\t")
		// dec := snafuToDecimal(line)
		// fmt.Print(dec)
		// fmt.Print("\t")
		// fmt.Println(decimalToSnafu(dec))
		// fmt.Println(d(line)
		sum += snafuToDecimal(line)
	}
	fmt.Println(sum)
	fmt.Println(decimalToSnafu(sum))

}

func Solve2(path string) {
	// data := util.GetFileLines(path)

	// for _, line := range data {

	// }
}

func snafuToDecimal(num string) int64 {
	size := len(num)
	var value int64
	for i := 0; i < size; i++ {
		digit := num[size-(i+1)]
		amp := int64(util.Power(5, i))
		switch digit {
		case 61:
			value += -2 * amp
		case 45:
			value += -1 * amp
		// case 48:
		// 	value += 0 * amp	unnessesary calculation
		case 49:
			value += 1 * amp
		case 50:
			value += 2 * amp
		}
	}
	return value
}

func decimalToSnafu(num int64) string {
	quinNum := decimalToQuintary(num)
	digitCount := len(quinNum)
	snafuNum := make([]rune, digitCount)

	for i := 0; i < digitCount; i++ {
		switch quinNum[i] {
		case 3:
			snafuNum[i] = '='
			quinNum[i+1]++
		case 4:
			snafuNum[i] = '-'
			quinNum[i+1]++
		case 5:
			snafuNum[i] = '0'
			quinNum[i+1]++
		case 6:
			snafuNum[i] = '0'
			quinNum[i+1]++
		default:
			val := strconv.Itoa(quinNum[i])
			snafuNum[i] = rune(val[0])
		}
	}

	// reverse array
	for i, j := 0, len(snafuNum)-1; i < j; i, j = i+1, j-1 {
		snafuNum[i], snafuNum[j] = snafuNum[j], snafuNum[i]
	}

	result := string(snafuNum)
	if result[:1] == "0" {
		result = result[1:]
	}

	return result
}

func decimalToQuintary(num int64) []int {
	if num < 5 {
		return []int{int(num), 0}
	}
	if num == 5 {
		return []int{0, 1, 0}
	}

	// var result int64
	var digits []int
	curr := num

	exponent := 1
	for curr > int64(util.Power(5, exponent)) {
		exponent++
	}

	digits = make([]int, exponent+1)
	exponent--

	for i := exponent; i >= 0; i-- {
		next := int64(util.Power(5, i))
		for curr >= next {
			curr -= next
			digits[i]++
		}
	}

	// for i := 0; i <= exponent; i++ {
	// 	result += int64(digits[i]) * int64(util.Power(10, i))
	// }

	return digits
}

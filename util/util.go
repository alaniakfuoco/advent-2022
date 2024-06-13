package util

import (
	"math/bits"
	"math/rand"
	"os"
	"strings"
)

func ReadFile(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return data
}

func GetFileLines(path string) []string {
	data := ReadFile(path)
	lines := strings.Split(string(data), "\r\n")
	return lines
}

func RandomString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func PermutationsUpTo(base, vals []string, num int) [][]string {
	result := make([][]string, 0)
	for i, s := range vals {
		next := append(base, s)
		result = append(result, next)
		if num > 1 {
			nnext := make([]string, len(next))
			nval := make([]string, len(vals))
			copy(nnext, next)
			copy(nval, vals)
			result = append(result, PermutationsUpTo(nnext, append(nval[:i], nval[i+1:]...), num-1)...)
		}
	}
	return result
}

func Permutations(base, vals []string, num int) [][]string {
	result := make([][]string, 0)
	for i, s := range vals {
		next := append(base, s)
		if num > 1 {
			nnext := make([]string, len(next))
			nval := make([]string, len(vals))
			copy(nnext, next)
			copy(nval, vals)
			result = append(result, PermutationsUpTo(nnext, append(nval[:i], nval[i+1:]...), num-1)...)
		} else {
			result = append(result, next)
		}
	}
	return result
}

func Combinations[T any](set []T, n int) (subsets [][]T) {
	length := uint(len(set))

	if n > len(set) {
		n = len(set)
	}

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		if n > 0 && bits.OnesCount(uint(subsetBits)) != n {
			continue
		}

		var subset []T

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset = append(subset, set[object])
			}
		}
		// add subset to subsets
		subsets = append(subsets, subset)
	}
	return subsets
}

func SplitStrArray(a []string, index int) []string {
	return append(a[:index], a[index+1:]...)
}

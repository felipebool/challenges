package main

import (
	"fmt"
	"math"
)

func getDifference(a string, b string) int {
	var diff int
	var deletions int

	hashA := make(map[string]int)
	hashB := make(map[string]int)

	for _, value := range a {
		hashA[string(value)]++
	}

	for _, value := range b {
		hashB[string(value)]++
	}

	for keyA, valueA := range hashA {
		valueB, existsB := hashB[keyA]

		if existsB {
			diff = int(math.Abs(float64(valueB - valueA)))
			deletions += diff
			continue
		}

		deletions += valueA
	}

	return deletions
}

func makeAnagram(a string, b string) int32 {
	deletions := getDifference(a, b)
	deletions += getDifference(b, a)

	return int32(deletions)
}

func main() {
	fmt.Println(makeAnagram("cde", "abc"))
}

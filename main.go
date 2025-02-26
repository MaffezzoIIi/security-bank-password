package main

import (
	"fmt"
	"math/rand/v2"
)

func permute(nums, current []int, used []bool, results *[][]int) {

	if len(current) == len(nums) {
		temp := make([]int, len(current))
		copy(temp, current)
		*results = append(*results, temp)
		return
	}

	for i := range nums {
		if !used[i] {
			used[i] = true
			current = append(current, nums[i])

			permute(nums, current, used, results)

			current = current[:len(current)-1]
			used[i] = false
		}
	}
}

func generatePairArray(results [][]int) [][]int {
	pn := make([][]int, 5)

	t := results[rand.IntN(len(results))]

	aux := 0
	for i := range 5 {
		pn[i] = make([]int, 2)
		pn[i][0] = t[aux]
		pn[i][1] = t[aux+1]
		aux += 2
	}

	return pn
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	var results [][]int

	used := make([]bool, len(numbers))

	permute(numbers, []int{}, used, &results)

	fmt.Println(generatePairArray(results))
}

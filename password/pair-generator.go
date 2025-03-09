package password

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

// aioefwsjrd

func GenerateValidatorDigit(pn [][]int) int {
	var validator int

	var fl []int

	for i := range pn {
		fl = append(fl, pn[i][0], pn[i][1])
	}

	key := "sioefwsjrd"
	count := 0
	for i := range fl {
		if (len(key) - 1) < count {
			count = 0
		}

		validator += (fl[i] * int(key[count]-'a'))

		count++
	}

	fmt.Println(validator)
	return validator % 99
}

func GeneratePermutions() [][]int {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	var results [][]int

	used := make([]bool, len(numbers))

	permute(numbers, []int{}, used, &results)

	return results
}

func GeneratePair(permutions [][]int) [][]int {
	return generatePairArray(permutions)
}

func ValidatePassword(permutions [][]int, password string) bool {
	if len(password) != 5 {
		return false
	}

	pn := make([][]int, 5)

	for i := range password {
		pn[i] = make([]int, 2)
		pn[i][0] = int(password[i] - '0')
		pn[i][1] = int(password[i+1] - '0')
		i++
	}

	return true
}

func contains(pn [][]int, pair []int) bool {
	for _, p := range pn {
		if p[0] == pair[0] && p[1] == pair[1] {
			return true
		}
	}
	return false
}

func ValidatePasswordPairs(pn [][]int, password [][]int) bool {
	for i := range password {
		pair := password[i]

		if !contains(pn, pair) {
			return false
		}
	}

	return true
}

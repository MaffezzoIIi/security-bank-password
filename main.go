package main

var limitBreak = 0

func permute(nums, current []int, used []bool, results *[][]int) {

	if len(current) == len(nums) {
		temp := make([]int, len(current))
		copy(temp, current)
		*results = append(*results, temp)
		return
	}

	if len(*results) == 10 {
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

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	var results [][]int

	used := make([]bool, len(numbers))

	permute(numbers, []int{}, used, &results)

}

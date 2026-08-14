package twosum

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for index, num := range nums {
		_, ok := seen[num]
		if ok {
			return []int{index, seen[num]}
		} else {
			seen[target-num] = index
		}
	}

	return []int{0, 0}
}

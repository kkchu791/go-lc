package longestconsecutivefolder

func longestConsecutive(nums []int) int {
	m := 0
	set := make(map[int]bool)

	for _, num := range nums {
		set[num] = true
	}

	for num := range set {
		if !set[num+1] {
			// start walking
			count := 1

			for set[num-1] {
				count++
				num--
			}

			m = max(m, count)
		}
	}

	return m
}

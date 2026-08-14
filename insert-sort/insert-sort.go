package insertsort

// 5, 2, 4, 6, 3

func InsertSort(a []int) []int {
	for j := 1; j < len(a); j++ {
		key := a[j]
		i := j - 1

		for i >= 0 && a[i] > key {
			a[i+1] = a[i]
			i--
		}

		a[i+1] = key

	}

	return a
}

// time: O(n^2) worst case, best: O(n)
// space O(1)

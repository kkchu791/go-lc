package parallelcoursesiii

import "slices"

func minimumTime(n int, relations [][]int, time []int) int {
	al := make(map[int][]int, 0)
	queue := make([]int, 0)
	inD := make([]int, n)
	maxTime := make([]int, n)

	for _, edge := range relations {
		al[edge[0]-1] = append(al[edge[0]-1], edge[1]-1)

		inD[edge[1]-1]++
	}

	for node, val := range inD {
		if val == 0 {
			queue = append(queue, node)
			maxTime[node] = time[node]
		}
	}

	for len(queue) > 0 {
		el := queue[0]
		queue = queue[1:]

		for _, neighbor := range al[el] {
			maxTime[neighbor] = max(maxTime[neighbor], maxTime[el]+time[neighbor])
			inD[neighbor]--

			if inD[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return slices.Max(maxTime)
}

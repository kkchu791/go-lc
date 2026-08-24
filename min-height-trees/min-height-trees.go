package minheighttrees

func findMinHeightTrees(n int, edges [][]int) []int {
	if len(edges) == 0 {
		return []int{0}
	}

	al := make(map[int][]int)
	inD := make([]int, n)

	for _, edge := range edges {
		al[edge[0]] = append(al[edge[0]], edge[1])
		al[edge[1]] = append(al[edge[1]], edge[0])
		inD[edge[0]]++
		inD[edge[1]]++
	}

	leaves := make([]int, 0)

	for i, degree := range inD {
		if degree == 1 {
			leaves = append(leaves, i)
		}
	}
	for n > 2 {
		size := len(leaves)
		n -= size

		for i := 0; i < size; i++ {
			leaf := leaves[i]
			inD[leaf]-- // shed leaf by decrementing him from inD

			// shed his neighbors next
			for _, neighbor := range al[leaf] {
				inD[neighbor]--

				// if this neighbor became a leaf add him to leaves
				if inD[neighbor] == 1 {
					leaves = append(leaves, neighbor)
				}
			}
		}
		leaves = leaves[size:]
	}

	return leaves

}

package validtree

func validTree(n int, edges [][]int) bool {
	al := make(map[int][]int, n)
	visit := make(map[int]struct{})

	for _, edge := range edges {
		al[edge[0]] = append(al[edge[0]], edge[1])
		al[edge[1]] = append(al[edge[1]], edge[0])
	}

	var hasCycle func(int, int) bool

	hasCycle = func(node, parent int) bool {
		if _, exists := visit[node]; exists {
			return true
		}

		visit[node] = struct{}{}

		for _, neighbor := range al[node] {
			if parent == neighbor {
				continue
			}

			if hasCycle(neighbor, node) {
				return true
			}
		}

		return false
	}

	if hasCycle(0, -1) {
		return false
	}

	return n == len(visit)
}

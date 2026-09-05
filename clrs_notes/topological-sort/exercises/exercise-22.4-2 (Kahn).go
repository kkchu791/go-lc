package exercises

func countSimplePaths() int {
	n := 4
	edges := [][]string{
		{"A", "D"},
		{"A", "B"},
		{"A", "C"},
		{"D", "B"},
		{"B", "C"},
	}

	source := "A"
	target := "C"

	stringRep := map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
		"D": 3,
	}

	inD := make([]int, n)
	queue := make([]string, 0)
	paths := make([]int, n)
	al := make(map[string][]string, 0)

	for _, edge := range edges {
		al[edge[0]] = append(al[edge[0]], edge[1])
		inD[stringRep[edge[1]]]++
	}

	paths[stringRep[source]] = 1

	intRep := map[int]string{
		0: "A",
		1: "B",
		2: "C",
		3: "D",
	}

	for idx, indegree := range inD {
		if indegree == 0 {
			queue = append(queue, intRep[idx])
		}
	}

	for len(queue) > 0 {
		el := queue[0]
		queue = queue[1:]

		for _, neigh := range al[el] {
			paths[stringRep[neigh]] += paths[stringRep[el]]

			inD[stringRep[neigh]]--

			if inD[stringRep[neigh]] == 0 {
				queue = append(queue, neigh)
			}
		}
	}

	return paths[stringRep[target]]
}

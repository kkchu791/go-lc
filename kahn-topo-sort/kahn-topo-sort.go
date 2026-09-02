package main

import "fmt"

func kahnTS(n int, edges [][]int) []int {

	al := make(map[int][]int, 0)
	transposeAl := make(map[int][]int, 0)
	inD := make([]int, n, n)
	queue := make([]int, 0)
	res := make([]int, 0)

	for _, edge := range edges {
		al[edge[0]] = append(al[edge[0]], edge[1])
		transposeAl[edge[1]] = append(transposeAl[edge[1]], edge[0])
	}

	for key, val := range transposeAl {
		inD[key] += len(val)
	}

	for node := range n {
		if _, exists := transposeAl[node]; !exists {
			queue = append(queue, node)
		}
	}

	for len(queue) > 0 {
		el := queue[0]
		queue = queue[1:]

		res = append(res, el)

		for _, neighbor := range al[el] {
			inD[neighbor]--

			if inD[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}

	}

	return res

}

func main() {
	n := 6
	edges := [][]int{
		{0, 1},
		{1, 2},
		{2, 3},
		{4, 5},
		{5, 1},
		{5, 2},
	}

	fmt.Println(kahnTS(n, edges))
}

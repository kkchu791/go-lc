package main

import (
	"fmt"
	"slices"
)

func main() {
	n := 3
	edges := [][]int{
		{0, 1},
		{1, 0},
		{1, 2},
	}
	fmt.Println(stronglyConnectedComponents(n, edges))
}

func stronglyConnectedComponents(n int, edges [][]int) []map[int]struct{} {
	og := orderGraph(n, edges)
	tg := transposeGraph(edges)
	return getStronglyConnectedNodes(tg, og)
}

func orderGraph(n int, edges [][]int) []int {
	visitedMap := make(map[int]struct{}, 0)
	visitedSlice := make([]int, 0)
	visiting := make(map[int]struct{}, 0)

	al := make(map[int][]int, 0)

	for _, edge := range edges {
		al[edge[0]] = append(al[edge[0]], edge[1])
	}

	var dfs func(int)

	dfs = func(node int) {
		if _, exists := visiting[node]; exists {
			return
		}

		if _, exists := visitedMap[node]; exists {
			return
		}

		visiting[node] = struct{}{}

		for _, neighbor := range al[node] {
			dfs(neighbor)
		}

		visitedMap[node] = struct{}{}
		visitedSlice = append(visitedSlice, node)
		delete(visiting, node)
	}

	for node := range n {
		dfs(node)
	}

	slices.Reverse(visitedSlice)
	return visitedSlice
}

func transposeGraph(edges [][]int) map[int][]int {
	tg := make(map[int][]int, 0)
	for _, edge := range edges {
		tg[edge[1]] = append(tg[edge[1]], edge[0])
	}
	return tg
}

func getStronglyConnectedNodes(tg map[int][]int, og []int) []map[int]struct{} {
	res := make([]map[int]struct{}, 0)
	visiting := make(map[int]struct{}, 0)
	visited := make(map[int]struct{}, 0)

	var dfs func(int)

	dfs = func(n int) {
		if _, exists := visiting[n]; exists {
			res = append(res, copyMap(visiting))
			return
		}

		if _, exists := visited[n]; exists {
			if len(visiting) > 0 {
				res = append(res, copyMap(visiting))

			}
			return
		}

		visiting[n] = struct{}{}

		for _, neighbor := range tg[n] {
			dfs(neighbor)
		}

		delete(visiting, n)
		visited[n] = struct{}{}
	}

	for _, n := range og {
		dfs(n)
	}

	return res
}

func copyMap(m map[int]struct{}) map[int]struct{} {
	set := make(map[int]struct{}, len(m))
	for key := range m {
		set[key] = struct{}{}
	}
	return set
}

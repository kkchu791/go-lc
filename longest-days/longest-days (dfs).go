package main

import "fmt"

type Node struct {
	Id   string
	Time int
}

func longestDays(nodes []Node, edges [][2]string) int {
	al := make(map[string][]string)
	res := 0
	time := make(map[string]int)
	memo := make(map[string]int)

	for _, node := range nodes {
		time[node.Id] = node.Time
	}

	for _, edge := range edges {
		al[edge[0]] = append(al[edge[0]], edge[1])
	}

	var dfs func(string) int
	dfs = func(task string) int {
		if savedMaxAtThisNode, exists := memo[task]; exists {
			return savedMaxAtThisNode
		}

		maxTimeAtThisNode := 0

		for _, neighbor := range al[task] {
			maxTimeAtThisNode = max(maxTimeAtThisNode, dfs(neighbor))
		}

		memo[task] = time[task] + maxTimeAtThisNode
		return memo[task]
	}

	for _, n := range nodes {
		res = max(res, dfs(n.Id))
	}

	return res
}

func main() {

	nodes := []Node{
		{"inventory_1", 3},
		{"inventory_2", 1},
		{"programming", 2},
		{"machine_1", 4},
		{"cutting_1", 4},
		{"cleaning", 7},
		{"shipping", 1},
	}

	edges := [][2]string{
		{"inventory_1", "inventory_2"},
		{"programming", "machine_1"},
		{"inventory_2", "machine_1"},
		{"machine_1", "cutting_1"},
		{"machine_1", "cleaning"},
		{"cleaning", "shipping"},
		{"cutting_1", "shipping"},
	}

	fmt.Println(longestDays(nodes, edges))
}

package main

import "fmt"

type Frame struct {
	node  int
	index int
}

type al map[int][]int

func hasCycleIterative(al map[int][]int, node int) bool {
	visiting := make(map[int]bool, 0)
	visited := make(map[int]bool, 0)

	stack := make([]Frame, 0)

	source := Frame{node: node, index: 0}
	stack = append(stack, source)

	for len(stack) > 0 {
		topIdx := len(stack) - 1
		cf := &stack[topIdx] //not possible to be the bug
		u := cf.node
		i := cf.index
		neighbors := al[u] //

		// we say something like, have we explore all the neighbors

		if i < len(neighbors) {
			v := al[u][i]
			cf.index++ // we're exploring so we increase the index to allow us to explore the next one later

			// we check if we've seen this node before
			if visiting[v] {
				return true
			}

			if visited[v] {
				fmt.Println(visited, "visited")
				return false
			}

			visiting[u] = true

			nn := Frame{node: v, index: 0}
			stack = append(stack, nn)
		} else {
			stack = stack[0:topIdx] //take him out
			delete(visiting, u)
			visited[u] = true
		}
	}

	fmt.Println(visited, "visited")

	return false
}

func main() {
	graphDAG := al{
		0: {1, 2},
		1: {3},
		2: {},
		3: {},
	}

	graphWithCycle := al{
		0: {1, 2},
		1: {3},
		2: {},
		3: {0},
	}

	graphWithHittingVisited := al{
		0: {1, 2},
		1: {3},
		2: {3},
		3: {},
	}

	fmt.Println(hasCycleIterative(graphDAG, 0), "answer")
	fmt.Println(hasCycleIterative(graphWithCycle, 0), "answer")
	fmt.Println(hasCycleIterative(graphWithHittingVisited, 0), "answer")
}

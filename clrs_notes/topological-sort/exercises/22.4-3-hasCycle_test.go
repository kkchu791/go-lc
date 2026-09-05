package exercises

import "testing"

func TestDetectCycle(t *testing.T) {
	n := 3
	edges := [][]int{
		{0, 1},
		{1, 2},
	}
	t.Log(detectCycle(n, edges))
}

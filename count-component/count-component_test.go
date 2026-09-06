package countcomponent

import (
	"fmt"
	"testing"
)

func TestCountComponent(t *testing.T) {
	n := 5
	edges := [][]int{
		{0, 1},
		{1, 2},
		{3, 4},
	}
	result := countComponents(n, edges)
	expected := 2

	fmt.Println(result, expected)

	if result == expected {
		fmt.Println("test passed")
	} else {
		t.Errorf("test did not pass")
	}
}

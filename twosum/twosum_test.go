package twosum

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	result := twoSum([]int{3, 2, 4}, 6)
	expected := []int{2, 1}

	fmt.Println(result[0], expected[0])

	if result[0] != expected[0] || result[1] != expected[1] {
		t.Errorf("hey test didn't pass")
	} else {
		fmt.Println("hey it passed")
	}
}

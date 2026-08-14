package insertsort

import (
	"fmt"
	"slices"
	"testing"
)

func TestInsertSort(t *testing.T) {
	result := InsertSort([]int{5, 2, 4, 6, 3})
	expected := []int{2, 3, 4, 5, 6}

	fmt.Println(result, expected)

	if slices.Equal(result, expected) {
		fmt.Println("test passed")
	} else {
		t.Errorf("test did not pass")
	}
}

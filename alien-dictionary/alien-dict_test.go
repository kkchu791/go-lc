package aliendictionary

import (
	"fmt"
	"testing"
)

func TestAlienDict_Valid(t *testing.T) {
	words := []string{"z", "o", "a"}
	result := alienDict(words)
	expected := ""

	fmt.Println(result, expected)

	if result == expected {
		fmt.Println("test passed")
	} else {
		t.Errorf("test did not pass")
	}
}

func TestAlienDict_Invalid(t *testing.T) {
	words := []string{"abc", "ab"}
	result := alienDict(words)
	expected := ""

	fmt.Println(result, expected)

	if result == expected {
		fmt.Println("test passed")
	} else {
		t.Errorf("test did not pass")
	}
}

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

func TestAlienDict_ZO(t *testing.T) {
	words := []string{"z", "o"}
	result := alienDict(words)
	expected := "zo"

	fmt.Println(result, expected)

	if result == expected {
		fmt.Println("test passed")
	} else {
		t.Errorf("test did not pass")
	}
}

func TestAlienDict_hernf(t *testing.T) {
	words := []string{"hrn", "hrf", "er", "enn", "rfnn"}
	result := alienDict(words)
	expected := "hernf"

	fmt.Println(result, expected)

	if result == expected {
		fmt.Println("test passed")
	} else {
		t.Errorf("test did not pass")
	}
}

func TestAlienDict_invalid_length(t *testing.T) {
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

func TestAlienDict_z(t *testing.T) {
	words := []string{"z"}
	result := alienDict(words)
	expected := "z"

	fmt.Println(result, expected)

	if result == expected {
		fmt.Println("test passed")
	} else {
		t.Errorf("test did not pass")
	}
}

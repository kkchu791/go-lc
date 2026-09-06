package aliendictionary

import (
	"fmt"
	"slices"
)

func alienDict(words []string) string {

	set := make(map[rune]struct{})
	al := make(map[rune][]rune)

	// because I loop with the expectation of 2 words i start at the second index
	// but for words having just 1 word, i still want to add it to the set
	// It'll just be vertices with no edges, no need for matching logic
	if len(words) == 1 {
		a := words[0]

		for _, r := range a {
			set[r] = struct{}{}
		}
	}

outer:
	for i := 1; i < len(words); i++ {

		a := words[i-1]
		b := words[i]

		for _, r := range a {
			set[r] = struct{}{}
		}

		for _, r := range b {
			set[r] = struct{}{}
		}

		runeA := []rune(a)
		runeB := []rune(b)

		j := 0

		for j < len(runeA) && j < len(runeB) {
			if runeA[j] != runeB[j] {
				al[runeA[j]] = append(al[runeA[j]], runeB[j])
				continue outer
			}

			j++
		}

		if len(runeA) > len(runeB) {
			return ""
		}
	}

	// Printing the set
	fmt.Print("set: map[")
	for k := range set {
		fmt.Printf("%c:{} ", k)
	}
	fmt.Println("]")

	//part II (topo sort the adj list)
	visited := make(map[rune]struct{})
	visitedSlice := make([]rune, 0)
	visiting := make(map[rune]struct{})

	var hasCycle func(rune) bool
	hasCycle = func(pc rune) bool {
		if _, exists := visited[pc]; exists {
			return false
		}

		if _, exists := visiting[pc]; exists {
			return true
		}

		visiting[pc] = struct{}{}

		for _, neigh := range al[pc] {
			if hasCycle(neigh) {
				return true
			}
		}

		visitedSlice = append(visitedSlice, pc)
		visited[pc] = struct{}{}
		delete(visiting, pc)

		return false
	}

	for pc := range set {
		if hasCycle(pc) {
			return ""
		}
	}

	slices.Reverse(visitedSlice)

	return string(visitedSlice)

}

// // Printing the set
// fmt.Print("set: map[")
// for k := range set {
// 	fmt.Printf("%c:{} ", k)
// }
// fmt.Println("]")

// // Printing the adjacency list (al)
// fmt.Print("al: map[")
// for k, values := range al {
// 	fmt.Printf("%c:[", k)
// 	for i, v := range values {
// 		fmt.Printf("%c", v)
// 		if i < len(values)-1 {
// 			fmt.Print(" ") // add a space between characters in the slice
// 		}
// 	}
// 	fmt.Print("] ")
// }
// fmt.Println("]")

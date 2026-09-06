package aliendictionary

import "fmt"

func alienDict(words []string) string {

	set := make(map[rune]struct{})
	al := make(map[rune][]rune)

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

	// Printing the adjacency list (al)
	fmt.Print("al: map[")
	for k, values := range al {
		fmt.Printf("%c:[", k)
		for i, v := range values {
			fmt.Printf("%c", v)
			if i < len(values)-1 {
				fmt.Print(" ") // add a space between characters in the slice
			}
		}
		fmt.Print("] ")
	}
	fmt.Println("]")
	return ""

}

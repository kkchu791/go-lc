package main

import "fmt"

func frequency(l []int) map[int]int {
	freq := make(map[int]int)

	for _, val := range l {
		_, exists := freq[val]

		if exists {
			freq[val]++
		} else {
			freq[val] = 1
		}
	}

	return freq
}

func main() {
	list := []int{1, 2, 3, 1}
	fmt.Println(frequency(list))
}

package main

import (
	"fmt"
	"slices"
	"strconv"
)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if !slices.Contains(wordList, endWord) {
		return 0
	}

	if !slices.Contains(wordList, beginWord) {
		wordList = append(wordList, beginWord)
	}

	al := make(map[string][]string, 0)
	for _, word := range wordList {
		al[word] = make([]string, 0)
	}

	for i := 0; i < len(wordList); i++ {
		for j := 0; j < len(wordList); j++ {
			word1 := wordList[i]
			word2 := wordList[j]

			count := 0
			k := 0

			for k < len(word1) {
				if word1[k] != word2[k] {
					count++
				}
				k++

				if count == 2 {
					break
				}
			}

			if count == 1 {
				al[word1] = append(al[word1], word2)
			}
		}
	}

	// part II BFS Approach
	// create a queue
	// pop from queu
	// explore neighbors
	// until you find the endword
	// add a secondary count helper with it

	queue := make([][]string, 0)
	visited := make(map[string]bool)
	source := []string{
		beginWord,
		"1",
	}

	queue = append(queue, source)

	for len(queue) > 0 {
		currEl := queue[0]
		queue = queue[1:]
		word := currEl[0]
		count := currEl[1]

		if visited[word] {
			continue
		}

		countNum, err := strconv.Atoi(count)
		if err != nil {
			fmt.Println("Error during conversion:", err)
			return 0
		}

		if word == endWord {
			return countNum
		}

		countNum = countNum + 1

		for _, neigh := range al[word] {

			str := strconv.Itoa(countNum)
			el := []string{
				neigh,
				str,
			}

			queue = append(queue, el)
		}

		visited[word] = true
	}

	return 0

}

func main() {

	beginWord := "hit"
	endWord := "cog"
	wordList := []string{
		"hot",
		"dot",
		"dog",
		"lot",
		"log",
		"cog",
	}
	fmt.Println(ladderLength(beginWord, endWord, wordList))
}

// part II DFS: Not working
// visited := make(map[string]int)
// visiting := make(map[string]bool)
// var getMinPath func(string) int

// getMinPath = func(word string) int {
// 	if visiting[word] {
// 		return 100
// 	}

// 	if val, exists := visited[word]; exists {
// 		return val
// 	}

// 	//base case
// 	if word == endWord {
// 		return 1
// 	}

// 	visiting[word] = true

// 	minPath := 100

// 	for _, neigh := range al[word] {
// 		minPath = min(minPath, getMinPath(neigh))
// 	}

// 	visited[word] = minPath
// 	delete(visiting, word)
// 	return minPath + 1
// }

// return getMinPath(beginWord)

// first attempt at getting words

// newSet := make(map[rune]struct{})
// 		for _, r := range currWord {
// 			newSet[r] = struct{}{}
// 		}

// 		for _, comparedWord := range wordList {
// 			if comparedWord == currWord {
// 				continue
// 			}

// 			count := 0
// 			// iterate through the word
// 			for _, r := range comparedWord {
// 				if _, exists := newSet[r]; !exists {
// 					count++
// 				}
// 			}

// 			// if count is 1 add to our adj list
// 			if count == 1 {
// 				al[currWord] = append(al[currWord], comparedWord)
// 			}
// 		}

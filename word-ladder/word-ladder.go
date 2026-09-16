package main

import "fmt"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	//Part 1: How do you tell 2 words are different by one letter

	// first check lengths match, if they don't it can't work
	// can add word1 chars to a set
	// iterate through word 2 and count how many letters don't match
	// if count ends up being 1, its a possible next Word

	//len(word1) != len(word2)
	//  immediately return 0
	// initialize letter diff
	//add word1 to a set
	// iterater through word2
	// if letter not in set
	// count++

	// if count == 1
	// add it to words 1 al

	wordList = append(wordList, beginWord)

	al := make(map[string][]string, 0)
	for _, word := range wordList {
		al[word] = make([]string, 0)
	}

	for _, currWord := range wordList {
		newSet := make(map[rune]struct{})
		for _, r := range currWord {
			newSet[r] = struct{}{}
		}

		for _, comparedWord := range wordList {
			if comparedWord == currWord {
				continue
			}

			count := 0
			// iterate through the word
			for _, r := range comparedWord {
				if _, exists := newSet[r]; !exists {
					count++
				}
			}

			// if count is 1 add to our adj list
			if count == 1 {
				al[currWord] = append(al[currWord], comparedWord)
			}
		}
	}

	fmt.Println(al, "al")
	fmt.Println(endWord, "endWord")

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

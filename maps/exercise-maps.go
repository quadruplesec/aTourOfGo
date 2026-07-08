package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)

	wordCounts := make(map[string]int)

	for i := 0; i < len(words); i++ {
		_, ok := wordCounts[words[i]]

		if !ok {
			wordCounts[words[i]] = 1
		} else {
			wordCounts[words[i]] += 1
		}
	}

	return wordCounts
}

func main() {
	wc.Test(WordCount)
}

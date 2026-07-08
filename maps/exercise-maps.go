package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)

	wordCounts := make(map[string]int)

	for _, word := range words { // Ranges make this easier to work with
        wordCounts[word]++ // No need to check if value is ok since Go defaults the value to 0
    }

    return wordCounts
}

func main() {
	wc.Test(WordCount)
}

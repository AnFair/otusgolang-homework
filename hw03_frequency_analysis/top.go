package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type Pair struct {
	word  string
	count int
}

func Top10(str string) []string {
	wordCountMap := make(map[string]int)
	words := strings.Fields(str)
	wordCountArr := []Pair{}
	result := []string{}
	for _, word := range words {
		wordCountMap[word]++
	}

	for word, count := range wordCountMap {
		wordCountArr = append(wordCountArr, Pair{word, count})
	}
	sort.Slice(wordCountArr, func(i, j int) bool {
		if wordCountArr[i].count == wordCountArr[j].count {
			return wordCountArr[i].word < wordCountArr[j].word
		}
		return wordCountArr[j].count < wordCountArr[i].count
	})

	for i, pair := range wordCountArr {
		if i == 10 {
			break
		}
		result = append(result, pair.word)
	}

	return result
}

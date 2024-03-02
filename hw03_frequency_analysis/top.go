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
	wordCountArr := make([]Pair, 0, 10)
	var result []string
	for _, word := range words {
		wordCountMap[word]++
	}

	for k, v := range wordCountMap {
		wordCountArr = append(wordCountArr, Pair{k, v})
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
		result[i] = pair.word
	}

	return result
}

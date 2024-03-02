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
	var wordCountMap = make(map[string]int)
	var words = strings.Fields(str)
	var wordCountArr []Pair
	var result []string
	for _, word := range words {
		if _, present := wordCountMap[word]; present {
			wordCountMap[word] = wordCountMap[word] + 1
		} else {
			wordCountMap[word] = 1
		}
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
		result = append(result, pair.word)
	}

	return result
}

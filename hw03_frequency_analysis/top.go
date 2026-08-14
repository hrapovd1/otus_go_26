package hw03frequencyanalysis

import (
	"cmp"
	"slices"
	"strings"
)

type WordCount struct {
	Value string
	Count int
}

func Top10(in string) []string {
	words := make(map[string]int)
	for w := range strings.FieldsSeq(in) {
		words[w]++
	}
	wordCounts := make([]WordCount, len(words))
	i := 0
	for w, c := range words {
		wordCounts[i].Value = w
		wordCounts[i].Count = c
		i++
	}
	slices.SortStableFunc(wordCounts, func(a, b WordCount) int {
		if firstCmp := cmp.Compare(b.Count, a.Count); firstCmp == 0 {
			return cmp.Compare(a.Value, b.Value)
		} else {
			return firstCmp
		}
	})
	out := make([]string, 0, 10)
	for i, wc := range wordCounts {
		if i > 9 {
			break
		}
		out = append(out, wc.Value)
	}

	return out
}

package src

import (
	"math/rand"
	"strings"
)

var Alphabet = []rune{'a', 'b'}

func CountAllEntry(str, s string) []int {
	var indexList []int
	for i := 0; i < len(str)-len(s)+1; i++ {
		relIndex := strings.HasPrefix(str[i:], s)
		if !relIndex {
			continue
		} else {
			indexList = append(indexList, i)
		}
	}
	return indexList
}

func ReplaceFromIndex(word, left, right string, index int) string {
	return word[:index] + strings.Replace(word[index:], left, right, 1)
}

func GenerateWords(minWordLen, maxWordLen int) string {
	var word strings.Builder
	lenght := rand.Intn(maxWordLen-minWordLen+1) + minWordLen
	for range lenght {
		word.WriteRune(Alphabet[rand.Intn(2)])
	}
	return word.String()
}

func ShortLex(a, b string) (string, string) {
	if len(a) == len(b) {
		for i, _ := range a {
			if a[i] < b[i] {
				return a, b
			} else if a[i] > b[i] {
				return b, a
			}
		}

	} else if len(a) > len(b) {
		return a, b
	} else {
		return b, a
	}
	return a, b
}

package src

import (
	"math/rand"
)

type RewritingRule struct {
	Left, Right string
}

func GenerateChain(minSteps, maxSteps int, word string, srs []RewritingRule) ([]string, string) {
	var chain []string
	steps := rand.Intn(maxSteps-minSteps+1) + minSteps
	currentWord := word
	chain = append(chain, currentWord)

	for range steps {
		orderRules := rand.Perm(len(srs))
		applied := false

		for _, k := range orderRules {
			indexCount := CountAllEntry(currentWord, srs[k].Left)
			if len(indexCount) != 0 {
				changes := indexCount[rand.Intn(len(indexCount))]
				currentWord = ReplaceFromIndex(currentWord, srs[k].Left, srs[k].Right, changes)
				chain = append(chain, currentWord)
				applied = true
				break
			}
		}
		if !applied {
			break
		}
	}

	return chain, currentWord
}

func AllVariantsToRewrite(wordStart string, srs_1 []RewritingRule) []string {
	var newStrings []string
	for _, rule := range srs_1 {
		ruleEntry := CountAllEntry(wordStart, rule.Left)
		if len(ruleEntry) > 0 {
			for j := 0; j < len(ruleEntry); j++ {
				newStrings = append(newStrings, ReplaceFromIndex(wordStart, rule.Left, rule.Right, ruleEntry[j]))
			}
		}
	}
	return newStrings
}

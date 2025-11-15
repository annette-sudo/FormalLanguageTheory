package transforms

import (
	"math/rand"
)

type RewritingRule struct {
	Left, Right string
}

var T = []RewritingRule{
	{Left: "abbaba", Right: "aabbbaa"},
	{Left: "aaa", Right: "abab"},
	{Left: "abba", Right: "baaab"},
	{Left: "bbbb", Right: "ba"},
}

var T_orient = []RewritingRule{
	{Left: "aabbbaa", Right: "abbaba"},
	{Left: "abab", Right: "aaa"},
	{Left: "baaab", Right: "abba"},
	{Left: "bbbb", Right: "ba"},
}

var T_1 = []RewritingRule{
	{Left: "bba", Right: "bab"},
	{Left: "aaaa", Right: "aaa"},
	{Left: "aaab", Right: "aaa"},
	{Left: "abaa", Right: "aaa"},
	{Left: "abab", Right: "aaa"},
	{Left: "baaa", Right: "aaa"},
	{Left: "baba", Right: "baab"},
	{Left: "bbbb", Right: "ba"},
	{Left: "baaba", Right: "aaa"},
	{Left: "baabb", Right: "aaa"},
	{Left: "babbb", Right: "baa"},
}

const MinWordLen = 10
const MinSteps = 10

const MaxWordLen = 30
const MaxSteps = 30

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

func RewritingRuleToString(srs RewritingRule) string {
	return srs.Left + " -> " + srs.Right
}
